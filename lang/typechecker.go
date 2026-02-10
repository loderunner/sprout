package lang

import "fmt"

type TypeError struct {
	Range   SourceRange
	Message string
}

func typeErrorf(expr Expr, format string, args ...any) *TypeError {
	return &TypeError{
		Range:   expr.Range(),
		Message: fmt.Sprintf(format, args...),
	}
}

func (err *TypeError) Error() string {
	return fmt.Sprintf("at %d:%d: %s", err.Range.StartLine, err.Range.StartCol, err.Message)
}

func TypeCheck(ast Expr) (Type, error) {
	ctx := NewContext()
	t, err := typeCheck(ctx, ast)
	if err != nil {
		return nil, err
	}
	return ctx.Subst.Apply(t), nil
}

func typeCheck(ctx *Context, expr Expr) (Type, error) {
	switch expr := expr.(type) {
	case IntLiteral:
		return checkIntLiteral(ctx, expr)
	case BoolLiteral:
		return checkBoolLiteral(ctx, expr)
	case VarExpr:
		return checkVarExpr(ctx, expr)
	case LetExpr:
		return checkLetExpr(ctx, expr)
	case IfExpr:
		return checkIfExpr(ctx, expr)
	case FunExpr:
		return checkFunExpr(ctx, expr)
	case AppExpr:
		return checkAppExpr(ctx, expr)
	case UnaryExpr:
		return checkUnaryExpr(ctx, expr)
	case BinaryExpr:
		return checkBinaryExpr(ctx, expr)
	default:
		panic(fmt.Sprintf("invalid expression: %T", expr))
	}
}

func checkIntLiteral(*Context, IntLiteral) (Type, error) {
	return IntType{}, nil
}

func checkBoolLiteral(*Context, BoolLiteral) (Type, error) {
	return BoolType{}, nil
}

func checkVarExpr(ctx *Context, expr VarExpr) (Type, error) {
	if varScheme, ok := ctx.Vars[expr.Name]; ok {
		return ctx.Instantiate(varScheme), nil
	}
	return nil, typeErrorf(expr, "unknown variable: %s", expr.Name)
}

func checkLetExpr(ctx *Context, expr LetExpr) (Type, error) {
	valueType, err := typeCheck(ctx, expr.Value)
	if err != nil {
		return nil, err
	}
	ctx = ctx.WithVar(expr.Name, ctx.Generalize(valueType))
	return typeCheck(ctx, expr.Body)
}

func checkIfExpr(ctx *Context, expr IfExpr) (Type, error) {
	condType, err := typeCheck(ctx, expr.Cond)
	if err != nil {
		return nil, err
	}
	if _, err = ctx.Subst.Unify(condType, BoolType{}); err != nil {
		return nil, typeErrorf(expr.Cond, "expected Bool condition, got %s", ctx.Subst.Apply(condType).TypeName())
	}
	thenType, err := typeCheck(ctx, expr.Then)
	if err != nil {
		return nil, err
	}
	elseType, err := typeCheck(ctx, expr.Else)
	if err != nil {
		return nil, err
	}
	if _, err = ctx.Subst.Unify(thenType, elseType); err != nil {
		return nil, typeErrorf(expr.Else, "type mismatch: expected %s, got %s", ctx.Subst.Apply(thenType).TypeName(), ctx.Subst.Apply(elseType).TypeName())
	}
	return thenType, nil
}

func checkFunExpr(ctx *Context, expr FunExpr) (Type, error) {
	var paramType Type
	if expr.ParamType == "" {
		// unannotated type
		paramType = ctx.NewTypeVar()
	} else {
		// annotated type
		var ok bool
		paramType, ok = ctx.Types[expr.ParamType]
		if !ok {
			return nil, typeErrorf(expr, "unknown type: %s", expr.ParamType)
		}
	}
	ctx = ctx.WithVar(expr.ParamName, Mono(paramType))
	returnType, err := typeCheck(ctx, expr.Body)
	if err != nil {
		return nil, err
	}
	paramType = ctx.Subst.Apply(paramType)
	return FunType{Param: paramType, Return: returnType}, nil
}

func checkAppExpr(ctx *Context, expr AppExpr) (Type, error) {
	funType, err := typeCheck(ctx, expr.Fun)
	if err != nil {
		return nil, err
	}
	funType = ctx.Subst.Apply(funType)

	argType, err := typeCheck(ctx, expr.Arg)
	if err != nil {
		return nil, err
	}
	argType = ctx.Subst.Apply(argType)

	switch funType := funType.(type) {
	case FunType:
		if _, err = ctx.Subst.Unify(argType, funType.Param); err != nil {
			return nil, typeErrorf(expr.Arg, "type mismatch: expected %s, got %s", ctx.Subst.Apply(funType.Param).TypeName(), argType.TypeName())
		}
		return funType.Return, nil
	case TypeVar:
		retType := ctx.NewTypeVar()
		if _, err = ctx.Subst.Unify(funType, FunType{Param: argType, Return: retType}); err != nil {
			// should never fail: `funType` is free, we're just binding it to argType -> ?retType
			return nil, typeErrorf(expr.Fun, "cannot apply as function")
		}
		return retType, nil
	default:
		return nil, typeErrorf(expr.Fun, "expected Function, got %s", funType.TypeName())
	}
}

func checkUnaryExpr(ctx *Context, expr UnaryExpr) (Type, error) {
	op := expr.Op

	operandType, err := typeCheck(ctx, expr.Expr)
	if err != nil {
		return nil, err
	}

	switch op {
	case "-":
		if _, err = ctx.Subst.Unify(operandType, IntType{}); err != nil {
			return nil, typeErrorf(expr.Expr, "expected Int, got %s", ctx.Subst.Apply(operandType).TypeName())
		}
		return IntType{}, nil
	case "!":
		if _, err = ctx.Subst.Unify(operandType, BoolType{}); err != nil {
			return nil, typeErrorf(expr.Expr, "expected Bool, got %s", ctx.Subst.Apply(operandType).TypeName())
		}
		return BoolType{}, nil
	default:
		panic(fmt.Sprintf("invalid unary operator: %s", op))
	}
}

func checkBinaryExpr(ctx *Context, expr BinaryExpr) (Type, error) {
	op := expr.Op

	leftType, err := typeCheck(ctx, expr.Left)
	if err != nil {
		return nil, err
	}

	rightType, err := typeCheck(ctx, expr.Right)
	if err != nil {
		return nil, err
	}

	switch op {
	case "+", "-", "*", "/":
		if _, err = ctx.Subst.Unify(leftType, IntType{}); err != nil {
			return nil, typeErrorf(expr.Left, "expected Int, got %s", ctx.Subst.Apply(leftType).TypeName())
		}
		if _, err = ctx.Subst.Unify(rightType, IntType{}); err != nil {
			return nil, typeErrorf(expr.Right, "expected Int, got %s", ctx.Subst.Apply(rightType).TypeName())
		}
		return IntType{}, nil
	case "<", "<=", ">", ">=":
		if _, err = ctx.Subst.Unify(leftType, IntType{}); err != nil {
			return nil, typeErrorf(expr.Left, "expected Int, got %s", ctx.Subst.Apply(leftType).TypeName())
		}
		if _, err = ctx.Subst.Unify(rightType, IntType{}); err != nil {
			return nil, typeErrorf(expr.Right, "expected Int, got %s", ctx.Subst.Apply(rightType).TypeName())
		}
		return BoolType{}, nil
	case "==", "!=":
		if _, err = ctx.Subst.Unify(leftType, rightType); err != nil {
			return nil, typeErrorf(expr, "type mismatch: %s and %s", ctx.Subst.Apply(leftType).TypeName(), ctx.Subst.Apply(rightType).TypeName())
		}
		return BoolType{}, nil
	case "&&", "||":
		if _, err = ctx.Subst.Unify(leftType, BoolType{}); err != nil {
			return nil, typeErrorf(expr.Left, "expected Bool, got %s", ctx.Subst.Apply(leftType).TypeName())
		}
		if _, err = ctx.Subst.Unify(rightType, BoolType{}); err != nil {
			return nil, typeErrorf(expr.Right, "expected Bool, got %s", ctx.Subst.Apply(rightType).TypeName())
		}
		return BoolType{}, nil
	default:
		panic(fmt.Sprintf("invalid binary operator: %s", op))
	}
}

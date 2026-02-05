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

func TypeCheck(ctx *Context, expr Expr) (Type, error) {
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
	if varType, ok := ctx.Vars[expr.Name]; ok {
		return varType, nil
	}
	return nil, typeErrorf(expr, "unknown variable: %s", expr.Name)
}

func checkLetExpr(ctx *Context, expr LetExpr) (Type, error) {
	valueType, err := TypeCheck(ctx, expr.Value)
	if err != nil {
		return nil, err
	}
	ctx = ctx.WithVar(expr.Name, valueType)
	return TypeCheck(ctx, expr.Body)
}

func checkIfExpr(ctx *Context, expr IfExpr) (Type, error) {
	condType, err := TypeCheck(ctx, expr.Cond)
	if err != nil {
		return nil, err
	}
	if _, ok := condType.(BoolType); !ok {
		return nil, typeErrorf(expr.Cond, "expected Bool condition, got %s", condType.TypeName())
	}
	thenType, err := TypeCheck(ctx, expr.Then)
	if err != nil {
		return nil, err
	}
	elseType, err := TypeCheck(ctx, expr.Else)
	if err != nil {
		return nil, err
	}
	if thenType != elseType {
		return nil, typeErrorf(expr.Else, "type mismatch: expected %s, got %s", thenType.TypeName(), elseType.TypeName())
	}
	return thenType, nil
}

func checkFunExpr(ctx *Context, expr FunExpr) (Type, error) {
	paramType, ok := ctx.Types[expr.ParamType]
	if !ok {
		return nil, typeErrorf(expr, "unknown type: %s", expr.ParamType)
	}
	ctx = ctx.WithVar(expr.ParamName, paramType)
	returnType, err := TypeCheck(ctx, expr.Body)
	if err != nil {
		return nil, err
	}
	return FunType{Param: paramType, Return: returnType}, nil
}

func checkAppExpr(ctx *Context, expr AppExpr) (Type, error) {
	funType, err := TypeCheck(ctx, expr.Fun)
	if err != nil {
		return nil, err
	}
	ft, ok := funType.(FunType)
	if !ok {
		return nil, typeErrorf(expr.Fun, "expected Function, got %s", funType.TypeName())
	}
	argType, err := TypeCheck(ctx, expr.Arg)
	if err != nil {
		return nil, err
	}
	if argType != ft.Param {
		return nil, typeErrorf(expr.Arg, "type mismatch: expected %s, got %s", ft.Param.TypeName(), argType.TypeName())
	}
	return ft.Return, nil
}

func checkUnaryExpr(ctx *Context, expr UnaryExpr) (Type, error) {
	op := expr.Op

	operandType, err := TypeCheck(ctx, expr.Expr)
	if err != nil {
		return nil, err
	}

	switch op {
	case "-":
		if _, ok := operandType.(IntType); !ok {
			return nil, typeErrorf(expr.Expr, "expected Int, got %s", operandType.TypeName())
		}
		return IntType{}, nil
	case "!":
		if _, ok := operandType.(BoolType); !ok {
			return nil, typeErrorf(expr.Expr, "expected Bool, got %s", operandType.TypeName())
		}
		return BoolType{}, nil
	default:
		panic(fmt.Sprintf("invalid unary operator: %s", op))
	}
}

func checkBinaryExpr(ctx *Context, expr BinaryExpr) (Type, error) {
	op := expr.Op

	leftType, err := TypeCheck(ctx, expr.Left)
	if err != nil {
		return nil, err
	}

	rightType, err := TypeCheck(ctx, expr.Right)
	if err != nil {
		return nil, err
	}

	switch op {
	case "+", "-", "*", "/":
		if _, ok := leftType.(IntType); !ok {
			return nil, typeErrorf(expr.Left, "expected Int, got %s", leftType.TypeName())
		}
		if _, ok := rightType.(IntType); !ok {
			return nil, typeErrorf(expr.Right, "expected Int, got %s", rightType.TypeName())
		}
		return IntType{}, nil
	case "<", "<=", ">", ">=":
		if _, ok := leftType.(IntType); !ok {
			return nil, typeErrorf(expr.Left, "expected Int, got %s", leftType.TypeName())
		}
		if _, ok := rightType.(IntType); !ok {
			return nil, typeErrorf(expr.Right, "expected Int, got %s", rightType.TypeName())
		}
		return BoolType{}, nil
	case "==", "!=":
		if leftType != rightType {
			return nil, typeErrorf(expr, "type mismatch: %s and %s", leftType.TypeName(), rightType.TypeName())
		}
		return BoolType{}, nil
	case "&&", "||":
		if _, ok := leftType.(BoolType); !ok {
			return nil, typeErrorf(expr.Left, "expected Bool, got %s", leftType.TypeName())
		}
		if _, ok := rightType.(BoolType); !ok {
			return nil, typeErrorf(expr.Right, "expected Bool, got %s", rightType.TypeName())
		}
		return BoolType{}, nil
	default:
		panic(fmt.Sprintf("invalid binary operator: %s", op))
	}
}

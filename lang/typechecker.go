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

func TypeCheck(ctx Context, expr Expr) (Type, error) {
	switch expr := expr.(type) {
	case IntLiteral:
		return IntType{}, nil
	case BoolLiteral:
		return BoolType{}, nil
	case VarExpr:
		if varType, ok := ctx[expr.Name]; ok {
			return varType, nil
		}
		return nil, typeErrorf(expr, "unknown variable: %s", expr.Name)
	case LetExpr:
		valueType, err := TypeCheck(ctx, expr.Value)
		if err != nil {
			return nil, err
		}
		ctx = ctx.Extend(expr.Name, valueType)
		return TypeCheck(ctx, expr.Body)
	case IfExpr:
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
	default:
		panic(fmt.Sprintf("unknown expression type: %T", expr))
	}
}

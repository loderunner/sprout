package lang

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/loderunner/sprout/lang/parser"
)

func rangeFromContext(ctx antlr.ParserRuleContext) SourceRange {
	startToken := ctx.GetStart()
	endToken := ctx.GetStop()
	return SourceRange{
		StartLine: uint(startToken.GetLine()),
		StartCol:  uint(startToken.GetColumn()) + 1,
		EndLine:   uint(endToken.GetLine()),
		EndCol:    uint(endToken.GetColumn()+len(endToken.GetText())) + 1,
	}
}

type ASTBuilder struct {
	*parser.BaseSproutVisitor
}

func NewASTBuilder() *ASTBuilder {
	return &ASTBuilder{}
}

func (b *ASTBuilder) VisitExpr(ctx *parser.ExprContext) Expr {
	if typeDefCtx := ctx.TypeDefExpr(); typeDefCtx != nil {
		return b.VisitTypeDefExpr(typeDefCtx.(*parser.TypeDefExprContext))
	}

	if letCtx := ctx.LetExpr(); letCtx != nil {
		return b.VisitLetExpr(letCtx.(*parser.LetExprContext))
	}

	if ifCtx := ctx.IfExpr(); ifCtx != nil {
		return b.VisitIfExpr(ifCtx.(*parser.IfExprContext))
	}

	if funCtx := ctx.FunExpr(); funCtx != nil {
		return b.VisitFunExpr(funCtx.(*parser.FunExprContext))
	}

	if compCtx := ctx.CompExpr(); compCtx != nil {
		return b.VisitCompExpr(compCtx.(*parser.CompExprContext))
	}

	panic(fmt.Sprintf("unexpected expression while parsing expression: `%s`", ctx.GetText()))
}

func (b *ASTBuilder) VisitTypeDefExpr(ctx *parser.TypeDefExprContext) Expr {
	ident := ctx.IDENT().GetText()
	typeExpr := b.VisitTypeExpr(ctx.TypeExpr().(*parser.TypeExprContext))
	body := b.VisitExpr(ctx.Expr().(*parser.ExprContext))
	return TypeDefExpr{
		Name:    ident,
		Type:    typeExpr,
		Body:    body,
		Located: Located{loc: rangeFromContext(ctx)},
	}
}

func (b *ASTBuilder) VisitLetExpr(ctx *parser.LetExprContext) Expr {
	name := ctx.IDENT().GetText()
	var typeExpr TypeExpr
	if typeCtx := ctx.TypeExpr(); typeCtx != nil {
		typeExpr = b.VisitTypeExpr(typeCtx.(*parser.TypeExprContext))
	}
	value := b.VisitExpr(ctx.Expr(0).(*parser.ExprContext))
	body := b.VisitExpr(ctx.Expr(1).(*parser.ExprContext))
	return LetExpr{
		Name:    name,
		Type:    typeExpr,
		Value:   value,
		Body:    body,
		Located: Located{loc: rangeFromContext(ctx)},
	}
}

func (b *ASTBuilder) VisitTypeExpr(ctx *parser.TypeExprContext) TypeExpr {
	primary := b.VisitPrimaryTypeExpr(ctx.PrimaryTypeExpr().(*parser.PrimaryTypeExprContext))
	if ctx.ARROW() == nil {
		return primary
	}
	ret := b.VisitTypeExpr(ctx.TypeExpr().(*parser.TypeExprContext))
	return ArrowTypeExpr{
		Param:   primary,
		Return:  ret,
		Located: Located{loc: rangeFromContext(ctx)},
	}
}

func (b *ASTBuilder) VisitPrimaryTypeExpr(ctx *parser.PrimaryTypeExprContext) TypeExpr {
	if ident := ctx.IDENT(); ident != nil {
		return NamedTypeExpr{
			Name:    ident.GetText(),
			Located: Located{loc: rangeFromContext(ctx)},
		}
	}
	return b.VisitTypeExpr(ctx.TypeExpr().(*parser.TypeExprContext))
}

func (b *ASTBuilder) VisitIfExpr(ctx *parser.IfExprContext) Expr {
	cond := b.VisitExpr(ctx.Expr(0).(*parser.ExprContext))
	then := b.VisitExpr(ctx.Expr(1).(*parser.ExprContext))
	els := b.VisitExpr(ctx.Expr(2).(*parser.ExprContext))
	return IfExpr{
		Cond:    cond,
		Then:    then,
		Else:    els,
		Located: Located{loc: rangeFromContext(ctx)},
	}
}

func (b *ASTBuilder) VisitFunExpr(ctx *parser.FunExprContext) Expr {
	paramName := ctx.IDENT().GetText()
	var paramType TypeExpr
	if typeCtx := ctx.TypeExpr(); typeCtx != nil {
		paramType = b.VisitTypeExpr(typeCtx.(*parser.TypeExprContext))
	}
	body := b.VisitExpr(ctx.Expr().(*parser.ExprContext))
	return FunExpr{
		ParamName: paramName,
		ParamType: paramType,
		Body:      body,
		Located:   Located{loc: rangeFromContext(ctx)},
	}
}

func (b *ASTBuilder) VisitCompExpr(ctx *parser.CompExprContext) Expr {
	left := b.VisitOrExpr(ctx.OrExpr(0).(*parser.OrExprContext))
	if opCtx := ctx.COMPOP(); opCtx != nil {
		op := opCtx.GetText()
		right := b.VisitOrExpr(ctx.OrExpr(1).(*parser.OrExprContext))
		return BinaryExpr{
			Op:      op,
			Left:    left,
			Right:   right,
			Located: Located{loc: rangeFromContext(ctx)},
		}
	}
	return left
}

func (b *ASTBuilder) VisitOrExpr(ctx *parser.OrExprContext) Expr {
	andExpr := b.VisitAndExpr(ctx.AndExpr().(*parser.AndExprContext))
	if orCtx := ctx.OrExpr(); orCtx != nil {
		left := b.VisitOrExpr(orCtx.(*parser.OrExprContext))
		return BinaryExpr{
			Op:      "||",
			Left:    left,
			Right:   andExpr,
			Located: Located{loc: rangeFromContext(ctx)},
		}
	}
	return andExpr
}

func (b *ASTBuilder) VisitAndExpr(ctx *parser.AndExprContext) Expr {
	termExpr := b.VisitTermExpr(ctx.TermExpr().(*parser.TermExprContext))
	if andCtx := ctx.AndExpr(); andCtx != nil {
		left := b.VisitAndExpr(andCtx.(*parser.AndExprContext))
		return BinaryExpr{
			Op:      "&&",
			Left:    left,
			Right:   termExpr,
			Located: Located{loc: rangeFromContext(ctx)},
		}
	}
	return termExpr
}

func (b *ASTBuilder) VisitTermExpr(ctx *parser.TermExprContext) Expr {
	factorExpr := b.VisitFactorExpr(ctx.FactorExpr().(*parser.FactorExprContext))
	if termCtx := ctx.TermExpr(); termCtx != nil {
		op := ctx.GetOp().GetText()
		left := b.VisitTermExpr(termCtx.(*parser.TermExprContext))
		return BinaryExpr{
			Op:      op,
			Left:    left,
			Right:   factorExpr,
			Located: Located{loc: rangeFromContext(ctx)},
		}
	}
	return factorExpr
}

func (b *ASTBuilder) VisitFactorExpr(ctx *parser.FactorExprContext) Expr {
	unaryExpr := b.VisitUnaryExpr(ctx.UnaryExpr().(*parser.UnaryExprContext))
	if factorCtx := ctx.FactorExpr(); factorCtx != nil {
		op := ctx.GetOp().GetText()
		left := b.VisitFactorExpr(factorCtx.(*parser.FactorExprContext))
		return BinaryExpr{
			Op:      op,
			Left:    left,
			Right:   unaryExpr,
			Located: Located{loc: rangeFromContext(ctx)},
		}
	}
	return unaryExpr
}

func (b *ASTBuilder) VisitUnaryExpr(ctx *parser.UnaryExprContext) Expr {
	if appCtx := ctx.AppExpr(); appCtx != nil {
		return b.VisitAppExpr(appCtx.(*parser.AppExprContext))
	}

	op := ctx.GetOp().GetText()
	expr := b.VisitUnaryExpr(ctx.UnaryExpr().(*parser.UnaryExprContext))
	return UnaryExpr{
		Op:      op,
		Expr:    expr,
		Located: Located{loc: rangeFromContext(ctx)},
	}
}

func (b *ASTBuilder) VisitAppExpr(ctx *parser.AppExprContext) Expr {
	if appCtx := ctx.AppExpr(); appCtx != nil {
		exprCtx := ctx.Expr()
		if exprCtx == nil {
			panic("missing expression in function application")
		}
		fun := b.VisitAppExpr(appCtx.(*parser.AppExprContext))
		arg := b.VisitExpr(exprCtx.(*parser.ExprContext))
		return AppExpr{
			Fun:     fun,
			Arg:     arg,
			Located: Located{loc: rangeFromContext(ctx)},
		}
	}

	if primaryCtx := ctx.PrimaryExpr(); primaryCtx != nil {
		return b.VisitPrimaryExpr(primaryCtx.(*parser.PrimaryExprContext))
	}

	panic(fmt.Sprintf("unexpected expression while parsing function application: `%s`", ctx.GetText()))
}

func (b *ASTBuilder) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) Expr {
	if parenCtx := ctx.Expr(); parenCtx != nil {
		return b.VisitExpr(parenCtx.(*parser.ExprContext))
	}

	if t := ctx.INT(); t != nil {
		val, err := strconv.Atoi(t.GetText())
		if err != nil {
			panic(fmt.Sprintf("failed to parse INT token `%q` at %d:%d: %s",
				t.GetText(),
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn()+1,
				err,
			))
		}
		return IntLiteral{
			Value:   val,
			Located: Located{loc: rangeFromContext(ctx)},
		}
	}

	if t := ctx.TRUE(); t != nil {
		return BoolLiteral{
			Value:   true,
			Located: Located{loc: rangeFromContext(ctx)},
		}
	}

	if t := ctx.FALSE(); t != nil {
		return BoolLiteral{
			Value:   false,
			Located: Located{loc: rangeFromContext(ctx)},
		}
	}

	if t := ctx.IDENT(); t != nil {
		return VarExpr{
			Name:    t.GetText(),
			Located: Located{loc: rangeFromContext(ctx)},
		}
	}

	panic(fmt.Sprintf("unexpected expression when parsing primary: `%s`", ctx.GetText()))
}

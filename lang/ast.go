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

func (b *ASTBuilder) VisitExpr(ctx *parser.ExprContext) any {
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

func (b *ASTBuilder) VisitLetExpr(ctx *parser.LetExprContext) any {
	name := ctx.IDENT().GetText()
	value := b.VisitExpr(ctx.Expr(0).(*parser.ExprContext)).(Expr)
	body := b.VisitExpr(ctx.Expr(1).(*parser.ExprContext)).(Expr)
	return LetExpr{
		Name:    name,
		Value:   value,
		Body:    body,
		located: located{loc: rangeFromContext(ctx)},
	}
}

func (b *ASTBuilder) VisitIfExpr(ctx *parser.IfExprContext) any {
	cond := b.VisitExpr(ctx.Expr(0).(*parser.ExprContext)).(Expr)
	then := b.VisitExpr(ctx.Expr(1).(*parser.ExprContext)).(Expr)
	els := b.VisitExpr(ctx.Expr(2).(*parser.ExprContext)).(Expr)
	return IfExpr{
		Cond:    cond,
		Then:    then,
		Else:    els,
		located: located{loc: rangeFromContext(ctx)},
	}
}

func (b *ASTBuilder) VisitFunExpr(ctx *parser.FunExprContext) any {
	paramName := ctx.IDENT(0).GetText()
	paramType := ctx.IDENT(1).GetText()
	body := b.VisitExpr(ctx.Expr().(*parser.ExprContext)).(Expr)
	return FunExpr{
		ParamName: paramName,
		ParamType: paramType,
		Body:      body,
		located:   located{loc: rangeFromContext(ctx)},
	}
}

func (b *ASTBuilder) VisitCompExpr(ctx *parser.CompExprContext) any {
	left := b.VisitOrExpr(ctx.OrExpr(0).(*parser.OrExprContext)).(Expr)
	if opCtx := ctx.COMPOP(); opCtx != nil {
		op := opCtx.GetText()
		right := b.VisitOrExpr(ctx.OrExpr(1).(*parser.OrExprContext)).(Expr)
		return BinaryExpr{
			Op:      op,
			Left:    left,
			Right:   right,
			located: located{loc: rangeFromContext(ctx)},
		}
	}
	return left
}

func (b *ASTBuilder) VisitOrExpr(ctx *parser.OrExprContext) any {
	andExpr := b.VisitAndExpr(ctx.AndExpr().(*parser.AndExprContext)).(Expr)
	if orCtx := ctx.OrExpr(); orCtx != nil {
		left := b.VisitOrExpr(orCtx.(*parser.OrExprContext)).(Expr)
		return BinaryExpr{
			Op:      "||",
			Left:    left,
			Right:   andExpr,
			located: located{loc: rangeFromContext(ctx)},
		}
	}
	return andExpr
}

func (b *ASTBuilder) VisitAndExpr(ctx *parser.AndExprContext) any {
	termExpr := b.VisitTermExpr(ctx.TermExpr().(*parser.TermExprContext)).(Expr)
	if andCtx := ctx.AndExpr(); andCtx != nil {
		left := b.VisitAndExpr(andCtx.(*parser.AndExprContext)).(Expr)
		return BinaryExpr{
			Op:      "&&",
			Left:    left,
			Right:   termExpr,
			located: located{loc: rangeFromContext(ctx)},
		}
	}
	return termExpr
}

func (b *ASTBuilder) VisitTermExpr(ctx *parser.TermExprContext) any {
	factorExpr := b.VisitFactorExpr(ctx.FactorExpr().(*parser.FactorExprContext)).(Expr)
	if termCtx := ctx.TermExpr(); termCtx != nil {
		op := ctx.GetOp().GetText()
		left := b.VisitTermExpr(termCtx.(*parser.TermExprContext)).(Expr)
		return BinaryExpr{
			Op:      op,
			Left:    left,
			Right:   factorExpr,
			located: located{loc: rangeFromContext(ctx)},
		}
	}
	return factorExpr
}

func (b *ASTBuilder) VisitFactorExpr(ctx *parser.FactorExprContext) any {
	unaryExpr := b.VisitUnaryExpr(ctx.UnaryExpr().(*parser.UnaryExprContext)).(Expr)
	if factorCtx := ctx.FactorExpr(); factorCtx != nil {
		op := ctx.GetOp().GetText()
		left := b.VisitFactorExpr(factorCtx.(*parser.FactorExprContext)).(Expr)
		return BinaryExpr{
			Op:      op,
			Left:    left,
			Right:   unaryExpr,
			located: located{loc: rangeFromContext(ctx)},
		}
	}
	return unaryExpr
}

func (b *ASTBuilder) VisitUnaryExpr(ctx *parser.UnaryExprContext) any {
	if appCtx := ctx.AppExpr(); appCtx != nil {
		return b.VisitAppExpr(appCtx.(*parser.AppExprContext)).(Expr)
	}

	op := ctx.GetOp().GetText()
	expr := b.VisitUnaryExpr(ctx.UnaryExpr().(*parser.UnaryExprContext)).(Expr)
	return UnaryExpr{
		Op:      op,
		Expr:    expr,
		located: located{loc: rangeFromContext(ctx)},
	}
}

func (b *ASTBuilder) VisitAppExpr(ctx *parser.AppExprContext) any {
	if appCtx := ctx.AppExpr(); appCtx != nil {
		exprCtx := ctx.Expr()
		if exprCtx == nil {
			panic("missing expression in function application")
		}
		fun := b.VisitAppExpr(appCtx.(*parser.AppExprContext)).(Expr)
		arg := b.VisitExpr(exprCtx.(*parser.ExprContext)).(Expr)
		return AppExpr{
			Fun:     fun,
			Arg:     arg,
			located: located{loc: rangeFromContext(ctx)},
		}
	}

	if primaryCtx := ctx.PrimaryExpr(); primaryCtx != nil {
		return b.VisitPrimaryExpr(primaryCtx.(*parser.PrimaryExprContext))
	}

	panic(fmt.Sprintf("unexpected expression while parsing function application: `%s`", ctx.GetText()))
}

func (b *ASTBuilder) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) any {
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
			located: located{loc: rangeFromContext(ctx)},
		}
	}

	if t := ctx.TRUE(); t != nil {
		return BoolLiteral{
			Value:   true,
			located: located{loc: rangeFromContext(ctx)},
		}
	}

	if t := ctx.FALSE(); t != nil {
		return BoolLiteral{
			Value:   false,
			located: located{loc: rangeFromContext(ctx)},
		}
	}

	if t := ctx.IDENT(); t != nil {
		return VarExpr{
			Name:    t.GetText(),
			located: located{loc: rangeFromContext(ctx)},
		}
	}

	panic(fmt.Sprintf("unexpected expression when parsing primary: `%s`", ctx.GetText()))
}

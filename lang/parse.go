package lang

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/loderunner/sprout/lang/parser"
)

type syntaxErrorListener struct {
	*antlr.DefaultErrorListener
	errors []*SyntaxError
}

type SyntaxError struct {
	Line    uint
	Col     uint
	Message string
}

func (err *SyntaxError) Error() string {
	return fmt.Sprintf("at %d:%d: %s", err.Line, err.Col, err.Message)
}

func (l *syntaxErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol any,
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {
	l.errors = append(l.errors, &SyntaxError{
		Line:    uint(line),
		Col:     uint(column + 1),
		Message: msg,
	})
}

func Parse(source string) (Expr, error) {
	errListener := syntaxErrorListener{}

	input := antlr.NewInputStream(string(source))
	lexer := parser.NewSproutLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&errListener)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSproutParser(tokens)
	p.RemoveErrorListeners()
	p.AddErrorListener(&errListener)

	tree := p.Expr()

	if len(errListener.errors) > 0 {
		return nil, errListener.errors[0]
	}

	builder := NewASTBuilder()
	expr := builder.VisitExpr(tree.(*parser.ExprContext)).(Expr)

	return expr, nil
}

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

	if appCtx := ctx.AppExpr(); appCtx != nil {
		return b.VisitAppExpr(appCtx.(*parser.AppExprContext))
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

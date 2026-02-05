package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/loderunner/sprout/lang"
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

func Parse(source string) (lang.Expr, error) {
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

	builder := lang.NewASTBuilder()
	expr := builder.VisitExpr(tree.(*parser.ExprContext)).(lang.Expr)

	return expr, nil
}

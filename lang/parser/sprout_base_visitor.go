// Code generated from Sprout.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Sprout
import "github.com/antlr4-go/antlr/v4"

type BaseSproutVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSproutVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSproutVisitor) VisitLetExpr(ctx *LetExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSproutVisitor) VisitIfExpr(ctx *IfExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSproutVisitor) VisitFunExpr(ctx *FunExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSproutVisitor) VisitAppExpr(ctx *AppExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSproutVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

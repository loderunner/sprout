// Code generated from Sprout.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Sprout
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SproutParser.
type SproutVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SproutParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by SproutParser#letExpr.
	VisitLetExpr(ctx *LetExprContext) interface{}

	// Visit a parse tree produced by SproutParser#ifExpr.
	VisitIfExpr(ctx *IfExprContext) interface{}

	// Visit a parse tree produced by SproutParser#funExpr.
	VisitFunExpr(ctx *FunExprContext) interface{}

	// Visit a parse tree produced by SproutParser#appExpr.
	VisitAppExpr(ctx *AppExprContext) interface{}

	// Visit a parse tree produced by SproutParser#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}
}

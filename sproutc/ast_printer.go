package main

import (
	"fmt"

	"github.com/loderunner/sprout/lang"
)

func printAST(ast lang.Expr) {
	printExpr(ast, "", true, true)
}

func printExpr(expr lang.Expr, prefix string, last bool, isRoot bool) {
	var nodePrefix string
	var childPrefix string

	if !isRoot {
		if last {
			nodePrefix = prefix + "└── "
			childPrefix = prefix + "    "
		} else {
			nodePrefix = prefix + "├── "
			childPrefix = prefix + "│   "
		}
	}

	switch e := expr.(type) {
	case lang.IntLiteral:
		fmt.Printf("%sIntLiteral(%d)\n", nodePrefix, e.Value)

	case lang.BoolLiteral:
		fmt.Printf("%sBoolLiteral(%t)\n", nodePrefix, e.Value)

	case lang.VarExpr:
		fmt.Printf("%sVarExpr(%s)\n", nodePrefix, e.Name)

	case lang.LetExpr:
		fmt.Printf("%sLetExpr(%s)\n", nodePrefix, e.Name)
		printExpr(e.Value, childPrefix, false, false)
		printExpr(e.Body, childPrefix, true, false)

	case lang.IfExpr:
		fmt.Printf("%sIfExpr\n", nodePrefix)
		printExpr(e.Cond, childPrefix, false, false)
		printExpr(e.Then, childPrefix, false, false)
		printExpr(e.Else, childPrefix, true, false)

	case lang.FunExpr:
		fmt.Printf("%sFunExpr(%s: %s)\n", nodePrefix, e.ParamName, e.ParamType)
		printExpr(e.Body, childPrefix, true, false)

	case lang.AppExpr:
		fmt.Printf("%sAppExpr\n", nodePrefix)
		printExpr(e.Fun, childPrefix, false, false)
		printExpr(e.Arg, childPrefix, true, false)

	case lang.UnaryExpr:
		fmt.Printf("%sUnaryExpr(%s)\n", nodePrefix, e.Op)
		printExpr(e.Expr, childPrefix, true, false)

	case lang.BinaryExpr:
		fmt.Printf("%sBinaryExpr(%s)\n", nodePrefix, e.Op)
		printExpr(e.Left, childPrefix, false, false)
		printExpr(e.Right, childPrefix, true, false)

	default:
		fmt.Printf("%sUnknownExpr\n", nodePrefix)
	}
}

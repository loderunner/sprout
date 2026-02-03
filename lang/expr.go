package lang

type SourceRange struct {
	StartLine uint
	StartCol  uint
	EndLine   uint
	EndCol    uint
}

type located struct {
	loc SourceRange
}

func (l located) Range() SourceRange {
	return l.loc
}

type Expr interface {
	exprNode()
	Range() SourceRange
}

type IntLiteral struct {
	located
	Value int
}

func (IntLiteral) exprNode() {}

type BoolLiteral struct {
	located
	Value bool
}

func (BoolLiteral) exprNode() {}

type VarExpr struct {
	located
	Name string
}

func (VarExpr) exprNode() {}

type LetExpr struct {
	located
	Name  string
	Value Expr
	Body  Expr
}

func (LetExpr) exprNode() {}

type IfExpr struct {
	located
	Cond Expr
	Then Expr
	Else Expr
}

func (IfExpr) exprNode() {}

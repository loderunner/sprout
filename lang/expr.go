package lang

type SourceRange struct {
	StartLine uint
	StartCol  uint
	EndLine   uint
	EndCol    uint
}

type Located struct {
	loc SourceRange
}

func (l Located) Range() SourceRange {
	return l.loc
}

type Ranger interface {
	Range() SourceRange
}

type Expr interface {
	exprNode()
	Ranger
}

type IntLiteral struct {
	Located
	Value int
}

func (IntLiteral) exprNode() {}

type BoolLiteral struct {
	Located
	Value bool
}

func (BoolLiteral) exprNode() {}

type VarExpr struct {
	Located
	Name string
}

func (VarExpr) exprNode() {}

type LetExpr struct {
	Located
	Name  string
	Type  TypeExpr
	Value Expr
	Body  Expr
}

func (LetExpr) exprNode() {}

type IfExpr struct {
	Located
	Cond Expr
	Then Expr
	Else Expr
}

func (IfExpr) exprNode() {}

type FunExpr struct {
	Located
	ParamName string
	ParamType TypeExpr
	Body      Expr
}

func (FunExpr) exprNode() {}

type AppExpr struct {
	Located
	Fun Expr
	Arg Expr
}

func (AppExpr) exprNode() {}

type UnaryExpr struct {
	Located
	Op   string
	Expr Expr
}

func (UnaryExpr) exprNode() {}

type BinaryExpr struct {
	Located
	Op    string
	Left  Expr
	Right Expr
}

func (BinaryExpr) exprNode() {}

type TypeExpr interface {
	typeExprNode()
	Ranger
}

type NamedTypeExpr struct {
	Located
	Name string
}

func (NamedTypeExpr) typeExprNode() {}

type ArrowTypeExpr struct {
	Located
	Param  TypeExpr
	Return TypeExpr
}

func (ArrowTypeExpr) typeExprNode() {}

package lang

type Type interface {
	TypeName() string
}

type IntType struct{}

func (IntType) TypeName() string {
	return "Int"
}

type BoolType struct{}

func (BoolType) TypeName() string {
	return "Bool"
}

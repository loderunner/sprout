package lang

import "fmt"

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

type FunType struct {
	Param  Type
	Return Type
}

func (f FunType) TypeName() string {
	return fmt.Sprintf("%s -> %s", f.Param.TypeName(), f.Return.TypeName())
}

type TypeVar struct {
	Id uint
}

func (t TypeVar) TypeName() string {
	return fmt.Sprintf("?T%d", t.Id)
}

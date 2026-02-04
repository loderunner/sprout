package lang

import "maps"

type Context struct {
	Vars  map[string]Type
	Types map[string]Type
}

var baseTypes = map[string]Type{
	"Bool": BoolType{},
	"Int":  IntType{},
}

func NewContext() *Context {
	return &Context{
		Vars:  make(map[string]Type),
		Types: baseTypes,
	}
}

func (c *Context) WithVar(name string, t Type) *Context {
	newCtx := &Context{
		Vars:  maps.Clone(c.Vars),
		Types: c.Types,
	}
	newCtx.Vars[name] = t
	return newCtx
}

func (c *Context) WithType(name string, t Type) *Context {
	newCtx := &Context{
		Vars:  c.Vars,
		Types: maps.Clone(c.Types),
	}
	newCtx.Types[name] = t
	return newCtx
}

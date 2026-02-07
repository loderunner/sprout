package lang

import "maps"

type Context struct {
	Vars   map[string]Type
	Types  map[string]Type
	Subst  Substitution
	nextTV *uint
}

var baseTypes = map[string]Type{
	"Bool": BoolType{},
	"Int":  IntType{},
}

func NewContext() *Context {
	var nextTV uint
	return &Context{
		Vars:   make(map[string]Type),
		Types:  baseTypes,
		Subst:  make(Substitution),
		nextTV: &nextTV,
	}
}

func (c *Context) clone() *Context {
	return &Context{
		Vars:   maps.Clone(c.Vars),
		Types:  maps.Clone(c.Types),
		Subst:  c.Subst,
		nextTV: c.nextTV,
	}
}

func (c *Context) WithVar(name string, t Type) *Context {
	newCtx := c.clone()
	newCtx.Vars[name] = t
	return newCtx
}

func (c *Context) WithType(name string, t Type) *Context {
	newCtx := c.clone()
	newCtx.Types[name] = t
	return newCtx
}

func (c *Context) NewTypeVar() TypeVar {
	tv := TypeVar{Id: *c.nextTV}
	*c.nextTV++
	return tv
}

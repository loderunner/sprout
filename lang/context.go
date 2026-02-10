package lang

import (
	"maps"
	"slices"
)

type Scheme struct {
	Bound []TypeVar
	Type  Type
}

func Mono(t Type) Scheme {
	return Scheme{Type: t}
}

type Context struct {
	Vars   map[string]Scheme
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
		Vars:   make(map[string]Scheme),
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

func (c *Context) WithVar(name string, s Scheme) *Context {
	newCtx := c.clone()
	newCtx.Vars[name] = s
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

func (c *Context) Instantiate(s Scheme) Type {
	subst := make(Substitution)
	for _, tv := range s.Bound {
		subst[tv] = c.NewTypeVar()
	}
	return c.Subst.Apply(subst.Apply(s.Type))
}

func (c *Context) Generalize(t Type) Scheme {
	t = c.Subst.Apply(t)

	tvs := make(map[TypeVar]struct{})
	findTypeVars(tvs, t)

	free := make(map[TypeVar]struct{})
	for _, s := range c.Vars {
		sType := c.Subst.Apply(s.Type)
		findTypeVars(free, sType)
		for _, bound := range s.Bound {
			delete(free, bound)
		}
	}

	for f := range maps.Keys(free) {
		delete(tvs, f)
	}

	return Scheme{
		Bound: slices.Collect(maps.Keys(tvs)),
		Type:  t,
	}
}

func findTypeVars(tvs map[TypeVar]struct{}, t Type) {
	switch t := t.(type) {
	case TypeVar:
		tvs[t] = struct{}{}
	case FunType:
		findTypeVars(tvs, t.Param)
		findTypeVars(tvs, t.Return)
		// do nothing for other types
	}
}

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
	localSubst := make(Substitution)
	for _, bound := range s.Bound {
		localSubst[bound] = c.NewTypeVar()
	}
	return c.Subst.Apply(localSubst.Apply(s.Type))
}

func (c *Context) Generalize(t Type) Scheme {
	t = c.Subst.Apply(t)

	// 1. Find type variables in a type
	tvs := make(map[TypeVar]struct{})
	findTypeVars(tvs, t)

	// 2. Find all free variables in the context
	free := make(map[TypeVar]struct{})
	for s := range maps.Values(c.Vars) {
		findTypeVars(free, s.Type)
		for _, bound := range s.Bound {
			delete(free, bound)
		}
	}

	// 3. Subtract context's free type variables
	for tv := range maps.Keys(free) {
		delete(tvs, tv)
	}

	return Scheme{
		Bound: slices.Collect(maps.Keys(tvs)),
		Type:  t,
	}
}

func findTypeVars(vars map[TypeVar]struct{}, t Type) {
	switch t := t.(type) {
	case TypeVar:
		vars[t] = struct{}{}
	case FunType:
		findTypeVars(vars, t.Param)
		findTypeVars(vars, t.Return)
	}
}

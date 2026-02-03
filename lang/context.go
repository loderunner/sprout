package lang

import "maps"

type Context map[string]Type

func (c Context) Extend(name string, t Type) Context {
	newCtx := maps.Clone(c)
	newCtx[name] = t
	return newCtx
}

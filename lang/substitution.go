package lang

import "fmt"

type Substitution map[TypeVar]Type

func (s Substitution) Apply(t Type) Type {
	switch t := t.(type) {
	case IntType, BoolType:
		return t
	case FunType:
		return FunType{Param: s.Apply(t.Param), Return: s.Apply(t.Return)}
	case TypeVar:
		if bound, ok := s[t]; ok {
			return s.Apply(bound)
		}
		return t
	default:
		panic(fmt.Sprintf("unknown type: %T", t))
	}
}

func (s Substitution) Unify(a, b Type) (Type, error) {
	a = s.Apply(a)
	b = s.Apply(b)

	if a == b {
		return a, nil
	}

	if a, ok := a.(TypeVar); ok {
		if s.Occurs(a, b) {
			return nil, fmt.Errorf("cannot unify %[1]s with %[2]s: %[1]s occurs in %[2]s", a.TypeName(), b.TypeName())
		}
		s[a] = b
		return b, nil
	}
	if b, ok := b.(TypeVar); ok {
		if s.Occurs(b, a) {
			return nil, fmt.Errorf("cannot unify %[1]s with %[2]s: %[1]s occurs in %[2]s", b.TypeName(), a.TypeName())
		}
		s[b] = a
		return a, nil
	}

	switch a := a.(type) {
	case BoolType:
		if _, ok := b.(BoolType); ok {
			return a, nil
		}
	case IntType:
		if _, ok := b.(IntType); ok {
			return b, nil
		}
	case FunType:
		if b, ok := b.(FunType); ok {
			p, err := s.Unify(a.Param, b.Param)
			if err != nil {
				return nil, err
			}
			r, err := s.Unify(a.Return, b.Return)
			if err != nil {
				return nil, err
			}
			return FunType{
				Param:  p,
				Return: r,
			}, nil
		}
	}

	return nil, fmt.Errorf("cannot unify %s with %s", a.TypeName(), b.TypeName())
}

func (s Substitution) Occurs(tv TypeVar, t Type) bool {
	t = s.Apply(t)
	switch t := t.(type) {
	case BoolType, IntType:
		return false
	case TypeVar:
		return tv == t
	case FunType:
		return s.Occurs(tv, t.Param) || s.Occurs(tv, t.Return)
	}
	panic(fmt.Sprintf("unknown type: %T", t))
}

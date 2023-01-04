// Code generated by mkunion. DO NOT EDIT.
package example

import (
	"github.com/widmogrod/mkunion/f"
)

type WherePredicateVisitor interface {
	VisitEq(v *Eq) any
	VisitAnd(v *And) any
	VisitOr(v *Or) any
	VisitPath(v *Path) any
}

type WherePredicate interface {
	Accept(g WherePredicateVisitor) any
}

func (r *Eq) Accept(v WherePredicateVisitor) any   { return v.VisitEq(r) }
func (r *And) Accept(v WherePredicateVisitor) any  { return v.VisitAnd(r) }
func (r *Or) Accept(v WherePredicateVisitor) any   { return v.VisitOr(r) }
func (r *Path) Accept(v WherePredicateVisitor) any { return v.VisitPath(r) }

var (
	_ WherePredicate = (*Eq)(nil)
	_ WherePredicate = (*And)(nil)
	_ WherePredicate = (*Or)(nil)
	_ WherePredicate = (*Path)(nil)
)

type WherePredicateOneOf struct {
	Eq   *Eq   `json:",omitempty"`
	And  *And  `json:",omitempty"`
	Or   *Or   `json:",omitempty"`
	Path *Path `json:",omitempty"`
}

func (r *WherePredicateOneOf) Accept(v WherePredicateVisitor) any {
	switch {
	case r.Eq != nil:
		return v.VisitEq(r.Eq)
	case r.And != nil:
		return v.VisitAnd(r.And)
	case r.Or != nil:
		return v.VisitOr(r.Or)
	case r.Path != nil:
		return v.VisitPath(r.Path)
	default:
		panic("unexpected")
	}
}

func (r *WherePredicateOneOf) Unwrap() WherePredicate {
	switch {
	case r.Eq != nil:
		return r.Eq
	case r.And != nil:
		return r.And
	case r.Or != nil:
		return r.Or
	case r.Path != nil:
		return r.Path
	}

	return nil
}

var _ WherePredicate = (*WherePredicateOneOf)(nil)

type mapWherePredicateToOneOf struct{}

func (t *mapWherePredicateToOneOf) VisitEq(v *Eq) any     { return &WherePredicateOneOf{Eq: v} }
func (t *mapWherePredicateToOneOf) VisitAnd(v *And) any   { return &WherePredicateOneOf{And: v} }
func (t *mapWherePredicateToOneOf) VisitOr(v *Or) any     { return &WherePredicateOneOf{Or: v} }
func (t *mapWherePredicateToOneOf) VisitPath(v *Path) any { return &WherePredicateOneOf{Path: v} }

var defaultMapWherePredicateToOneOf WherePredicateVisitor = &mapWherePredicateToOneOf{}

func MapWherePredicateToOneOf(v WherePredicate) *WherePredicateOneOf {
	return v.Accept(defaultMapWherePredicateToOneOf).(*WherePredicateOneOf)
}

func MustMatchWherePredicate[TOut any](
	x WherePredicate,
	f1 func(x *Eq) TOut,
	f2 func(x *And) TOut,
	f3 func(x *Or) TOut,
	f4 func(x *Path) TOut,
) TOut {
	return f.MustMatch4(x, f1, f2, f3, f4)
}

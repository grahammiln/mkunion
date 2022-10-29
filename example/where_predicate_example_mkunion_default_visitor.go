// Code generated by mkunion. DO NOT EDIT.
package example

type WherePredicateDefaultVisitor[A any] struct {
	Default A
	OnEq    func(x *Eq) A
	OnAnd   func(x *And) A
	OnOr    func(x *Or) A
	OnPath  func(x *Path) A
}

func (t *WherePredicateDefaultVisitor[A]) VisitEq(v *Eq) any {
	if t.OnEq != nil {
		return t.OnEq(v)
	}
	return t.Default
}
func (t *WherePredicateDefaultVisitor[A]) VisitAnd(v *And) any {
	if t.OnAnd != nil {
		return t.OnAnd(v)
	}
	return t.Default
}
func (t *WherePredicateDefaultVisitor[A]) VisitOr(v *Or) any {
	if t.OnOr != nil {
		return t.OnOr(v)
	}
	return t.Default
}
func (t *WherePredicateDefaultVisitor[A]) VisitPath(v *Path) any {
	if t.OnPath != nil {
		return t.OnPath(v)
	}
	return t.Default
}
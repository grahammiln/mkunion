// Code generated by mkunion. DO NOT EDIT.
package shape

type ShapeDefaultVisitor[A any] struct {
	Default       A
	OnAny         func(x *Any) A
	OnRefName     func(x *RefName) A
	OnBooleanLike func(x *BooleanLike) A
	OnStringLike  func(x *StringLike) A
	OnNumberLike  func(x *NumberLike) A
	OnListLike    func(x *ListLike) A
	OnMapLike     func(x *MapLike) A
	OnStructLike  func(x *StructLike) A
	OnUnionLike   func(x *UnionLike) A
}

func (t *ShapeDefaultVisitor[A]) VisitAny(v *Any) any {
	if t.OnAny != nil {
		return t.OnAny(v)
	}
	return t.Default
}
func (t *ShapeDefaultVisitor[A]) VisitRefName(v *RefName) any {
	if t.OnRefName != nil {
		return t.OnRefName(v)
	}
	return t.Default
}
func (t *ShapeDefaultVisitor[A]) VisitBooleanLike(v *BooleanLike) any {
	if t.OnBooleanLike != nil {
		return t.OnBooleanLike(v)
	}
	return t.Default
}
func (t *ShapeDefaultVisitor[A]) VisitStringLike(v *StringLike) any {
	if t.OnStringLike != nil {
		return t.OnStringLike(v)
	}
	return t.Default
}
func (t *ShapeDefaultVisitor[A]) VisitNumberLike(v *NumberLike) any {
	if t.OnNumberLike != nil {
		return t.OnNumberLike(v)
	}
	return t.Default
}
func (t *ShapeDefaultVisitor[A]) VisitListLike(v *ListLike) any {
	if t.OnListLike != nil {
		return t.OnListLike(v)
	}
	return t.Default
}
func (t *ShapeDefaultVisitor[A]) VisitMapLike(v *MapLike) any {
	if t.OnMapLike != nil {
		return t.OnMapLike(v)
	}
	return t.Default
}
func (t *ShapeDefaultVisitor[A]) VisitStructLike(v *StructLike) any {
	if t.OnStructLike != nil {
		return t.OnStructLike(v)
	}
	return t.Default
}
func (t *ShapeDefaultVisitor[A]) VisitUnionLike(v *UnionLike) any {
	if t.OnUnionLike != nil {
		return t.OnUnionLike(v)
	}
	return t.Default
}

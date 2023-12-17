// Code generated by mkunion. DO NOT EDIT.
package shape

import "github.com/widmogrod/mkunion/f"
import "github.com/widmogrod/mkunion/x/shared"
import "encoding/json"
import "fmt"

//mkunion-extension:visitor

type ShapeVisitor interface {
	VisitAny(v *Any) any
	VisitRefName(v *RefName) any
	VisitAliasLike(v *AliasLike) any
	VisitBooleanLike(v *BooleanLike) any
	VisitStringLike(v *StringLike) any
	VisitNumberLike(v *NumberLike) any
	VisitListLike(v *ListLike) any
	VisitMapLike(v *MapLike) any
	VisitStructLike(v *StructLike) any
	VisitUnionLike(v *UnionLike) any
}

type Shape interface {
	AcceptShape(g ShapeVisitor) any
}

func (r *Any) AcceptShape(v ShapeVisitor) any         { return v.VisitAny(r) }
func (r *RefName) AcceptShape(v ShapeVisitor) any     { return v.VisitRefName(r) }
func (r *AliasLike) AcceptShape(v ShapeVisitor) any   { return v.VisitAliasLike(r) }
func (r *BooleanLike) AcceptShape(v ShapeVisitor) any { return v.VisitBooleanLike(r) }
func (r *StringLike) AcceptShape(v ShapeVisitor) any  { return v.VisitStringLike(r) }
func (r *NumberLike) AcceptShape(v ShapeVisitor) any  { return v.VisitNumberLike(r) }
func (r *ListLike) AcceptShape(v ShapeVisitor) any    { return v.VisitListLike(r) }
func (r *MapLike) AcceptShape(v ShapeVisitor) any     { return v.VisitMapLike(r) }
func (r *StructLike) AcceptShape(v ShapeVisitor) any  { return v.VisitStructLike(r) }
func (r *UnionLike) AcceptShape(v ShapeVisitor) any   { return v.VisitUnionLike(r) }

var (
	_ Shape = (*Any)(nil)
	_ Shape = (*RefName)(nil)
	_ Shape = (*AliasLike)(nil)
	_ Shape = (*BooleanLike)(nil)
	_ Shape = (*StringLike)(nil)
	_ Shape = (*NumberLike)(nil)
	_ Shape = (*ListLike)(nil)
	_ Shape = (*MapLike)(nil)
	_ Shape = (*StructLike)(nil)
	_ Shape = (*UnionLike)(nil)
)

func MatchShape[TOut any](
	x Shape,
	f1 func(x *Any) TOut,
	f2 func(x *RefName) TOut,
	f3 func(x *AliasLike) TOut,
	f4 func(x *BooleanLike) TOut,
	f5 func(x *StringLike) TOut,
	f6 func(x *NumberLike) TOut,
	f7 func(x *ListLike) TOut,
	f8 func(x *MapLike) TOut,
	f9 func(x *StructLike) TOut,
	f10 func(x *UnionLike) TOut,
	df func(x Shape) TOut,
) TOut {
	return f.Match10(x, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, df)
}

func MatchShapeR2[TOut1, TOut2 any](
	x Shape,
	f1 func(x *Any) (TOut1, TOut2),
	f2 func(x *RefName) (TOut1, TOut2),
	f3 func(x *AliasLike) (TOut1, TOut2),
	f4 func(x *BooleanLike) (TOut1, TOut2),
	f5 func(x *StringLike) (TOut1, TOut2),
	f6 func(x *NumberLike) (TOut1, TOut2),
	f7 func(x *ListLike) (TOut1, TOut2),
	f8 func(x *MapLike) (TOut1, TOut2),
	f9 func(x *StructLike) (TOut1, TOut2),
	f10 func(x *UnionLike) (TOut1, TOut2),
	df func(x Shape) (TOut1, TOut2),
) (TOut1, TOut2) {
	return f.Match10R2(x, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, df)
}

func MustMatchShape[TOut any](
	x Shape,
	f1 func(x *Any) TOut,
	f2 func(x *RefName) TOut,
	f3 func(x *AliasLike) TOut,
	f4 func(x *BooleanLike) TOut,
	f5 func(x *StringLike) TOut,
	f6 func(x *NumberLike) TOut,
	f7 func(x *ListLike) TOut,
	f8 func(x *MapLike) TOut,
	f9 func(x *StructLike) TOut,
	f10 func(x *UnionLike) TOut,
) TOut {
	return f.MustMatch10(x, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10)
}

func MustMatchShapeR0(
	x Shape,
	f1 func(x *Any),
	f2 func(x *RefName),
	f3 func(x *AliasLike),
	f4 func(x *BooleanLike),
	f5 func(x *StringLike),
	f6 func(x *NumberLike),
	f7 func(x *ListLike),
	f8 func(x *MapLike),
	f9 func(x *StructLike),
	f10 func(x *UnionLike),
) {
	f.MustMatch10R0(x, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10)
}

func MustMatchShapeR2[TOut1, TOut2 any](
	x Shape,
	f1 func(x *Any) (TOut1, TOut2),
	f2 func(x *RefName) (TOut1, TOut2),
	f3 func(x *AliasLike) (TOut1, TOut2),
	f4 func(x *BooleanLike) (TOut1, TOut2),
	f5 func(x *StringLike) (TOut1, TOut2),
	f6 func(x *NumberLike) (TOut1, TOut2),
	f7 func(x *ListLike) (TOut1, TOut2),
	f8 func(x *MapLike) (TOut1, TOut2),
	f9 func(x *StructLike) (TOut1, TOut2),
	f10 func(x *UnionLike) (TOut1, TOut2),
) (TOut1, TOut2) {
	return f.MustMatch10R2(x, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10)
}

// mkunion-extension:shape
func ShapeShape() Shape {
	return &UnionLike{
		Name:          "Shape",
		PkgName:       "shape",
		PkgImportName: "github.com/widmogrod/mkunion/x/shape",
		Variant: []Shape{
			AnyShape(),
			RefNameShape(),
			AliasLikeShape(),
			BooleanLikeShape(),
			StringLikeShape(),
			NumberLikeShape(),
			ListLikeShape(),
			MapLikeShape(),
			StructLikeShape(),
			UnionLikeShape(),
		},
	}
}

func AnyShape() Shape {
	return &StructLike{
		Name:          "Any",
		PkgName:       "shape",
		PkgImportName: "github.com/widmogrod/mkunion/x/shape",
	}
}

func RefNameShape() Shape {
	return &StructLike{
		Name:          "RefName",
		PkgName:       "shape",
		PkgImportName: "github.com/widmogrod/mkunion/x/shape",
		Fields: []*FieldLike{
			{
				Name: "Name",
				Type: &StringLike{},
			},
			{
				Name: "PkgName",
				Type: &StringLike{},
			},
			{
				Name: "pkgImportName",
				Type: &StringLike{},
			},
			{
				Name: "IsPointer",
				Type: &BooleanLike{},
			},
			{
				Name: "Indexed",
				Type: &ListLike{
					Element: &RefName{
						Name:          "Shape",
						PkgName:       "shape",
						PkgImportName: "github.com/widmogrod/mkunion/x/shape",
						IsPointer:     false,
					},
					ElementIsPointer: false,
				},
			},
		},
	}
}

func AliasLikeShape() Shape {
	return &StructLike{
		Name:          "AliasLike",
		PkgName:       "shape",
		PkgImportName: "github.com/widmogrod/mkunion/x/shape",
		Fields: []*FieldLike{
			{
				Name: "Name",
				Type: &StringLike{},
			},
			{
				Name: "PkgName",
				Type: &StringLike{},
			},
			{
				Name: "pkgImportName",
				Type: &StringLike{},
			},
			{
				Name: "IsAlias",
				Type: &BooleanLike{},
			},
			{
				Name: "Type",
				Type: &RefName{
					Name:          "Shape",
					PkgName:       "shape",
					PkgImportName: "github.com/widmogrod/mkunion/x/shape",
					IsPointer:     false,
				},
			},
		},
	}
}

func BooleanLikeShape() Shape {
	return &StructLike{
		Name:          "BooleanLike",
		PkgName:       "shape",
		PkgImportName: "github.com/widmogrod/mkunion/x/shape",
	}
}

func StringLikeShape() Shape {
	return &StructLike{
		Name:          "StringLike",
		PkgName:       "shape",
		PkgImportName: "github.com/widmogrod/mkunion/x/shape",
	}
}

func NumberLikeShape() Shape {
	return &StructLike{
		Name:          "NumberLike",
		PkgName:       "shape",
		PkgImportName: "github.com/widmogrod/mkunion/x/shape",
		Fields: []*FieldLike{
			{
				Name: "Kind",
				Type: &RefName{
					Name:          "NumberKind",
					PkgName:       "shape",
					PkgImportName: "github.com/widmogrod/mkunion/x/shape",
					IsPointer:     false,
				},
			},
		},
	}
}

func ListLikeShape() Shape {
	return &StructLike{
		Name:          "ListLike",
		PkgName:       "shape",
		PkgImportName: "github.com/widmogrod/mkunion/x/shape",
		Fields: []*FieldLike{
			{
				Name: "Element",
				Type: &RefName{
					Name:          "Shape",
					PkgName:       "shape",
					PkgImportName: "github.com/widmogrod/mkunion/x/shape",
					IsPointer:     false,
				},
			},
			{
				Name: "ElementIsPointer",
				Type: &BooleanLike{},
			},
			{
				Name: "ArrayLen",
				Type: &NumberLike{},
			},
		},
	}
}

func MapLikeShape() Shape {
	return &StructLike{
		Name:          "MapLike",
		PkgName:       "shape",
		PkgImportName: "github.com/widmogrod/mkunion/x/shape",
		Fields: []*FieldLike{
			{
				Name: "Key",
				Type: &RefName{
					Name:          "Shape",
					PkgName:       "shape",
					PkgImportName: "github.com/widmogrod/mkunion/x/shape",
					IsPointer:     false,
				},
			},
			{
				Name: "Val",
				Type: &RefName{
					Name:          "Shape",
					PkgName:       "shape",
					PkgImportName: "github.com/widmogrod/mkunion/x/shape",
					IsPointer:     false,
				},
			},
			{
				Name: "KeyIsPointer",
				Type: &BooleanLike{},
			},
			{
				Name: "ValIsPointer",
				Type: &BooleanLike{},
			},
		},
	}
}

func StructLikeShape() Shape {
	return &StructLike{
		Name:          "StructLike",
		PkgName:       "shape",
		PkgImportName: "github.com/widmogrod/mkunion/x/shape",
		Fields: []*FieldLike{
			{
				Name: "Name",
				Type: &StringLike{},
			},
			{
				Name: "PkgName",
				Type: &StringLike{},
			},
			{
				Name: "pkgImportName",
				Type: &StringLike{},
			},
			{
				Name: "TypeParams",
				Type: &ListLike{
					Element: &RefName{
						Name:          "TypeParam",
						PkgName:       "shape",
						PkgImportName: "github.com/widmogrod/mkunion/x/shape",
						IsPointer:     false,
					},
					ElementIsPointer: false,
				},
			},
			{
				Name: "Fields",
				Type: &ListLike{
					Element: &RefName{
						Name:          "FieldLike",
						PkgName:       "shape",
						PkgImportName: "github.com/widmogrod/mkunion/x/shape",
						IsPointer:     true,
					},
					ElementIsPointer: true,
				},
			},
		},
	}
}

func UnionLikeShape() Shape {
	return &StructLike{
		Name:          "UnionLike",
		PkgName:       "shape",
		PkgImportName: "github.com/widmogrod/mkunion/x/shape",
		Fields: []*FieldLike{
			{
				Name: "Name",
				Type: &StringLike{},
			},
			{
				Name: "PkgName",
				Type: &StringLike{},
			},
			{
				Name: "pkgImportName",
				Type: &StringLike{},
			},
			{
				Name: "Variant",
				Type: &ListLike{
					Element: &RefName{
						Name:          "Shape",
						PkgName:       "shape",
						PkgImportName: "github.com/widmogrod/mkunion/x/shape",
						IsPointer:     false,
					},
					ElementIsPointer: false,
				},
			},
		},
	}
}

// mkunion-extension:json
type ShapeUnionJSON struct {
	Type        string          `json:"$type,omitempty"`
	Any         json.RawMessage `json:"shape.Any,omitempty"`
	RefName     json.RawMessage `json:"shape.RefName,omitempty"`
	AliasLike   json.RawMessage `json:"shape.AliasLike,omitempty"`
	BooleanLike json.RawMessage `json:"shape.BooleanLike,omitempty"`
	StringLike  json.RawMessage `json:"shape.StringLike,omitempty"`
	NumberLike  json.RawMessage `json:"shape.NumberLike,omitempty"`
	ListLike    json.RawMessage `json:"shape.ListLike,omitempty"`
	MapLike     json.RawMessage `json:"shape.MapLike,omitempty"`
	StructLike  json.RawMessage `json:"shape.StructLike,omitempty"`
	UnionLike   json.RawMessage `json:"shape.UnionLike,omitempty"`
}

func ShapeFromJSON(x []byte) (Shape, error) {
	var data ShapeUnionJSON
	err := json.Unmarshal(x, &data)
	if err != nil {
		return nil, err
	}

	switch data.Type {
	case "shape.Any":
		return AnyFromJSON(data.Any)
	case "shape.RefName":
		return RefNameFromJSON(data.RefName)
	case "shape.AliasLike":
		return AliasLikeFromJSON(data.AliasLike)
	case "shape.BooleanLike":
		return BooleanLikeFromJSON(data.BooleanLike)
	case "shape.StringLike":
		return StringLikeFromJSON(data.StringLike)
	case "shape.NumberLike":
		return NumberLikeFromJSON(data.NumberLike)
	case "shape.ListLike":
		return ListLikeFromJSON(data.ListLike)
	case "shape.MapLike":
		return MapLikeFromJSON(data.MapLike)
	case "shape.StructLike":
		return StructLikeFromJSON(data.StructLike)
	case "shape.UnionLike":
		return UnionLikeFromJSON(data.UnionLike)
	}

	if data.Any != nil {
		return AnyFromJSON(data.Any)
	} else if data.RefName != nil {
		return RefNameFromJSON(data.RefName)
	} else if data.AliasLike != nil {
		return AliasLikeFromJSON(data.AliasLike)
	} else if data.BooleanLike != nil {
		return BooleanLikeFromJSON(data.BooleanLike)
	} else if data.StringLike != nil {
		return StringLikeFromJSON(data.StringLike)
	} else if data.NumberLike != nil {
		return NumberLikeFromJSON(data.NumberLike)
	} else if data.ListLike != nil {
		return ListLikeFromJSON(data.ListLike)
	} else if data.MapLike != nil {
		return MapLikeFromJSON(data.MapLike)
	} else if data.StructLike != nil {
		return StructLikeFromJSON(data.StructLike)
	} else if data.UnionLike != nil {
		return UnionLikeFromJSON(data.UnionLike)
	}

	return nil, fmt.Errorf("unknown type %s", data.Type)
}

func ShapeToJSON(x Shape) ([]byte, error) {
	if x == nil {
		return nil, nil
	}
	return MustMatchShapeR2(
		x,
		func(x *Any) ([]byte, error) {
			body, err := AnyToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(ShapeUnionJSON{
				Type: "shape.Any",
				Any:  body,
			})
		},
		func(x *RefName) ([]byte, error) {
			body, err := RefNameToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(ShapeUnionJSON{
				Type:    "shape.RefName",
				RefName: body,
			})
		},
		func(x *AliasLike) ([]byte, error) {
			body, err := AliasLikeToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(ShapeUnionJSON{
				Type:      "shape.AliasLike",
				AliasLike: body,
			})
		},
		func(x *BooleanLike) ([]byte, error) {
			body, err := BooleanLikeToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(ShapeUnionJSON{
				Type:        "shape.BooleanLike",
				BooleanLike: body,
			})
		},
		func(x *StringLike) ([]byte, error) {
			body, err := StringLikeToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(ShapeUnionJSON{
				Type:       "shape.StringLike",
				StringLike: body,
			})
		},
		func(x *NumberLike) ([]byte, error) {
			body, err := NumberLikeToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(ShapeUnionJSON{
				Type:       "shape.NumberLike",
				NumberLike: body,
			})
		},
		func(x *ListLike) ([]byte, error) {
			body, err := ListLikeToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(ShapeUnionJSON{
				Type:     "shape.ListLike",
				ListLike: body,
			})
		},
		func(x *MapLike) ([]byte, error) {
			body, err := MapLikeToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(ShapeUnionJSON{
				Type:    "shape.MapLike",
				MapLike: body,
			})
		},
		func(x *StructLike) ([]byte, error) {
			body, err := StructLikeToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(ShapeUnionJSON{
				Type:       "shape.StructLike",
				StructLike: body,
			})
		},
		func(x *UnionLike) ([]byte, error) {
			body, err := UnionLikeToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(ShapeUnionJSON{
				Type:      "shape.UnionLike",
				UnionLike: body,
			})
		},
	)
}

func AnyFromJSON(x []byte) (*Any, error) {
	var result *Any = new(Any)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		}

		return fmt.Errorf("shape.AnyFromJSON: unknown key %s", key)
	})

	return result, err
}

func AnyToJSON(x *Any) ([]byte, error) {
	return json.Marshal(map[string]json.RawMessage{})
}
func (self *Any) MarshalJSON() ([]byte, error) {
	return AnyToJSON(self)
}

func (self *Any) UnmarshalJSON(x []byte) error {
	n, err := AnyFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func RefNameFromJSON(x []byte) (*RefName, error) {
	var result *RefName = new(RefName)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		case "Name":
			return json.Unmarshal(value, &result.Name)
		case "PkgName":
			return json.Unmarshal(value, &result.PkgName)
		case "pkgImportName":
			return json.Unmarshal(value, &result.PkgImportName)
		case "IsPointer":
			return json.Unmarshal(value, &result.IsPointer)
		case "Indexed":
			res, err := shared.JSONToListWithDeserializer(value, result.Indexed, ShapeFromJSON)
			if err != nil {
				return fmt.Errorf("shape._FromJSON: field Shape %w", err)
			}
			result.Indexed = res
			return nil
		}

		return fmt.Errorf("shape.RefNameFromJSON: unknown key %s", key)
	})

	return result, err
}

func RefNameToJSON(x *RefName) ([]byte, error) {
	field_Name, err := json.Marshal(x.Name)
	if err != nil {
		return nil, err
	}
	field_PkgName, err := json.Marshal(x.PkgName)
	if err != nil {
		return nil, err
	}
	field_PkgImportName, err := json.Marshal(x.PkgImportName)
	if err != nil {
		return nil, err
	}
	field_IsPointer, err := json.Marshal(x.IsPointer)
	if err != nil {
		return nil, err
	}
	field_Indexed, err := shared.JSONListFromSerializer(x.Indexed, ShapeToJSON)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"Name":          field_Name,
		"PkgName":       field_PkgName,
		"pkgImportName": field_PkgImportName,
		"IsPointer":     field_IsPointer,
		"Indexed":       field_Indexed,
	})
}
func (self *RefName) MarshalJSON() ([]byte, error) {
	return RefNameToJSON(self)
}

func (self *RefName) UnmarshalJSON(x []byte) error {
	n, err := RefNameFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func AliasLikeFromJSON(x []byte) (*AliasLike, error) {
	var result *AliasLike = new(AliasLike)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		case "Name":
			return json.Unmarshal(value, &result.Name)
		case "PkgName":
			return json.Unmarshal(value, &result.PkgName)
		case "pkgImportName":
			return json.Unmarshal(value, &result.PkgImportName)
		case "IsAlias":
			return json.Unmarshal(value, &result.IsAlias)
		case "Type":
			res, err := ShapeFromJSON(value)
			if err != nil {
				return fmt.Errorf("shape._FromJSON: field Shape %w", err)
			}
			result.Type = res
			return nil
		}

		return fmt.Errorf("shape.AliasLikeFromJSON: unknown key %s", key)
	})

	return result, err
}

func AliasLikeToJSON(x *AliasLike) ([]byte, error) {
	field_Name, err := json.Marshal(x.Name)
	if err != nil {
		return nil, err
	}
	field_PkgName, err := json.Marshal(x.PkgName)
	if err != nil {
		return nil, err
	}
	field_PkgImportName, err := json.Marshal(x.PkgImportName)
	if err != nil {
		return nil, err
	}
	field_IsAlias, err := json.Marshal(x.IsAlias)
	if err != nil {
		return nil, err
	}
	field_Type, err := ShapeToJSON(x.Type)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"Name":          field_Name,
		"PkgName":       field_PkgName,
		"pkgImportName": field_PkgImportName,
		"IsAlias":       field_IsAlias,
		"Type":          field_Type,
	})
}
func (self *AliasLike) MarshalJSON() ([]byte, error) {
	return AliasLikeToJSON(self)
}

func (self *AliasLike) UnmarshalJSON(x []byte) error {
	n, err := AliasLikeFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func BooleanLikeFromJSON(x []byte) (*BooleanLike, error) {
	var result *BooleanLike = new(BooleanLike)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		}

		return fmt.Errorf("shape.BooleanLikeFromJSON: unknown key %s", key)
	})

	return result, err
}

func BooleanLikeToJSON(x *BooleanLike) ([]byte, error) {
	return json.Marshal(map[string]json.RawMessage{})
}
func (self *BooleanLike) MarshalJSON() ([]byte, error) {
	return BooleanLikeToJSON(self)
}

func (self *BooleanLike) UnmarshalJSON(x []byte) error {
	n, err := BooleanLikeFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func StringLikeFromJSON(x []byte) (*StringLike, error) {
	var result *StringLike = new(StringLike)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		}

		return fmt.Errorf("shape.StringLikeFromJSON: unknown key %s", key)
	})

	return result, err
}

func StringLikeToJSON(x *StringLike) ([]byte, error) {
	return json.Marshal(map[string]json.RawMessage{})
}
func (self *StringLike) MarshalJSON() ([]byte, error) {
	return StringLikeToJSON(self)
}

func (self *StringLike) UnmarshalJSON(x []byte) error {
	n, err := StringLikeFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func NumberLikeFromJSON(x []byte) (*NumberLike, error) {
	var result *NumberLike = new(NumberLike)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		case "Kind":
			res, err := NumberKindFromJSON(value)
			if err != nil {
				return fmt.Errorf("shape._FromJSON: field NumberKind %w", err)
			}
			result.Kind = res
			return nil
		}

		return fmt.Errorf("shape.NumberLikeFromJSON: unknown key %s", key)
	})

	return result, err
}

func NumberLikeToJSON(x *NumberLike) ([]byte, error) {
	field_Kind, err := NumberKindToJSON(x.Kind)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"Kind": field_Kind,
	})
}
func (self *NumberLike) MarshalJSON() ([]byte, error) {
	return NumberLikeToJSON(self)
}

func (self *NumberLike) UnmarshalJSON(x []byte) error {
	n, err := NumberLikeFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func ListLikeFromJSON(x []byte) (*ListLike, error) {
	var result *ListLike = new(ListLike)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		case "Element":
			res, err := ShapeFromJSON(value)
			if err != nil {
				return fmt.Errorf("shape._FromJSON: field Shape %w", err)
			}
			result.Element = res
			return nil
		case "ElementIsPointer":
			return json.Unmarshal(value, &result.ElementIsPointer)
		case "ArrayLen":
			return json.Unmarshal(value, &result.ArrayLen)
		}

		return fmt.Errorf("shape.ListLikeFromJSON: unknown key %s", key)
	})

	return result, err
}

func ListLikeToJSON(x *ListLike) ([]byte, error) {
	field_Element, err := ShapeToJSON(x.Element)
	if err != nil {
		return nil, err
	}
	field_ElementIsPointer, err := json.Marshal(x.ElementIsPointer)
	if err != nil {
		return nil, err
	}
	field_ArrayLen, err := json.Marshal(x.ArrayLen)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"Element":          field_Element,
		"ElementIsPointer": field_ElementIsPointer,
		"ArrayLen":         field_ArrayLen,
	})
}
func (self *ListLike) MarshalJSON() ([]byte, error) {
	return ListLikeToJSON(self)
}

func (self *ListLike) UnmarshalJSON(x []byte) error {
	n, err := ListLikeFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func MapLikeFromJSON(x []byte) (*MapLike, error) {
	var result *MapLike = new(MapLike)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		case "Key":
			res, err := ShapeFromJSON(value)
			if err != nil {
				return fmt.Errorf("shape._FromJSON: field Shape %w", err)
			}
			result.Key = res
			return nil
		case "Val":
			res, err := ShapeFromJSON(value)
			if err != nil {
				return fmt.Errorf("shape._FromJSON: field Shape %w", err)
			}
			result.Val = res
			return nil
		case "KeyIsPointer":
			return json.Unmarshal(value, &result.KeyIsPointer)
		case "ValIsPointer":
			return json.Unmarshal(value, &result.ValIsPointer)
		}

		return fmt.Errorf("shape.MapLikeFromJSON: unknown key %s", key)
	})

	return result, err
}

func MapLikeToJSON(x *MapLike) ([]byte, error) {
	field_Key, err := ShapeToJSON(x.Key)
	if err != nil {
		return nil, err
	}
	field_Val, err := ShapeToJSON(x.Val)
	if err != nil {
		return nil, err
	}
	field_KeyIsPointer, err := json.Marshal(x.KeyIsPointer)
	if err != nil {
		return nil, err
	}
	field_ValIsPointer, err := json.Marshal(x.ValIsPointer)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"Key":          field_Key,
		"Val":          field_Val,
		"KeyIsPointer": field_KeyIsPointer,
		"ValIsPointer": field_ValIsPointer,
	})
}
func (self *MapLike) MarshalJSON() ([]byte, error) {
	return MapLikeToJSON(self)
}

func (self *MapLike) UnmarshalJSON(x []byte) error {
	n, err := MapLikeFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func StructLikeFromJSON(x []byte) (*StructLike, error) {
	var result *StructLike = new(StructLike)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		case "Name":
			return json.Unmarshal(value, &result.Name)
		case "PkgName":
			return json.Unmarshal(value, &result.PkgName)
		case "pkgImportName":
			return json.Unmarshal(value, &result.PkgImportName)
		case "TypeParams":
			return json.Unmarshal(value, &result.TypeParams)
		case "Fields":
			return json.Unmarshal(value, &result.Fields)
		}

		return fmt.Errorf("shape.StructLikeFromJSON: unknown key %s", key)
	})

	return result, err
}

func StructLikeToJSON(x *StructLike) ([]byte, error) {
	field_Name, err := json.Marshal(x.Name)
	if err != nil {
		return nil, err
	}
	field_PkgName, err := json.Marshal(x.PkgName)
	if err != nil {
		return nil, err
	}
	field_PkgImportName, err := json.Marshal(x.PkgImportName)
	if err != nil {
		return nil, err
	}
	field_TypeParams, err := json.Marshal(x.TypeParams)
	if err != nil {
		return nil, err
	}
	field_Fields, err := json.Marshal(x.Fields)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"Name":          field_Name,
		"PkgName":       field_PkgName,
		"pkgImportName": field_PkgImportName,
		"TypeParams":    field_TypeParams,
		"Fields":        field_Fields,
	})
}
func (self *StructLike) MarshalJSON() ([]byte, error) {
	return StructLikeToJSON(self)
}

func (self *StructLike) UnmarshalJSON(x []byte) error {
	n, err := StructLikeFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func UnionLikeFromJSON(x []byte) (*UnionLike, error) {
	var result *UnionLike = new(UnionLike)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		case "Name":
			return json.Unmarshal(value, &result.Name)
		case "PkgName":
			return json.Unmarshal(value, &result.PkgName)
		case "pkgImportName":
			return json.Unmarshal(value, &result.PkgImportName)
		case "Variant":
			res, err := shared.JSONToListWithDeserializer(value, result.Variant, ShapeFromJSON)
			if err != nil {
				return fmt.Errorf("shape._FromJSON: field Shape %w", err)
			}
			result.Variant = res
			return nil
		}

		return fmt.Errorf("shape.UnionLikeFromJSON: unknown key %s", key)
	})

	return result, err
}

func UnionLikeToJSON(x *UnionLike) ([]byte, error) {
	field_Name, err := json.Marshal(x.Name)
	if err != nil {
		return nil, err
	}
	field_PkgName, err := json.Marshal(x.PkgName)
	if err != nil {
		return nil, err
	}
	field_PkgImportName, err := json.Marshal(x.PkgImportName)
	if err != nil {
		return nil, err
	}
	field_Variant, err := shared.JSONListFromSerializer(x.Variant, ShapeToJSON)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"Name":          field_Name,
		"PkgName":       field_PkgName,
		"pkgImportName": field_PkgImportName,
		"Variant":       field_Variant,
	})
}
func (self *UnionLike) MarshalJSON() ([]byte, error) {
	return UnionLikeToJSON(self)
}

func (self *UnionLike) UnmarshalJSON(x []byte) error {
	n, err := UnionLikeFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

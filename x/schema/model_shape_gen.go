// Code generated by mkunion. DO NOT EDIT.
package schema

import (
	"github.com/widmogrod/mkunion/x/shape"
)

func init() {
	shape.Register(SchemaShape())
	shape.Register(NoneShape())
	shape.Register(BoolShape())
	shape.Register(NumberShape())
	shape.Register(StringShape())
	shape.Register(BinaryShape())
	shape.Register(ListShape())
	shape.Register(MapShape())
	shape.Register(FieldShape())
}

//shape:shape

func SchemaShape() shape.Shape {
	return &shape.UnionLike{
		Name:          "Schema",
		PkgName:       "schema",
		PkgImportName: "github.com/widmogrod/mkunion/x/schema",
		Variant: []shape.Shape{
			NoneShape(),
			BoolShape(),
			NumberShape(),
			StringShape(),
			BinaryShape(),
			ListShape(),
			MapShape(),
		},
	}
}

func NoneShape() shape.Shape {
	return &shape.StructLike{
		Name:          "None",
		PkgName:       "schema",
		PkgImportName: "github.com/widmogrod/mkunion/x/schema",
	}
}

func BoolShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "Bool",
		PkgName:       "schema",
		PkgImportName: "github.com/widmogrod/mkunion/x/schema",
		Type:          &shape.PrimitiveLike{Kind: &shape.BooleanLike{}},
	}
}

func NumberShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "Number",
		PkgName:       "schema",
		PkgImportName: "github.com/widmogrod/mkunion/x/schema",
		Type: &shape.PrimitiveLike{
			Kind: &shape.NumberLike{
				Kind: &shape.Float64{},
			},
		},
	}
}

func StringShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "String",
		PkgName:       "schema",
		PkgImportName: "github.com/widmogrod/mkunion/x/schema",
		Type:          &shape.PrimitiveLike{Kind: &shape.StringLike{}},
	}
}

func BinaryShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "Binary",
		PkgName:       "schema",
		PkgImportName: "github.com/widmogrod/mkunion/x/schema",
		Type: &shape.ListLike{
			Element: &shape.PrimitiveLike{
				Kind: &shape.NumberLike{
					Kind: &shape.UInt8{},
				},
			},
		},
	}
}

func ListShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "List",
		PkgName:       "schema",
		PkgImportName: "github.com/widmogrod/mkunion/x/schema",
		Type: &shape.ListLike{
			Element: &shape.RefName{
				Name:          "Schema",
				PkgName:       "schema",
				PkgImportName: "github.com/widmogrod/mkunion/x/schema",
			},
		},
	}
}

func MapShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "Map",
		PkgName:       "schema",
		PkgImportName: "github.com/widmogrod/mkunion/x/schema",
		Type: &shape.MapLike{
			Key: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
			Val: &shape.RefName{
				Name:          "Schema",
				PkgName:       "schema",
				PkgImportName: "github.com/widmogrod/mkunion/x/schema",
			},
		},
	}
}

//shape:shape
func FieldShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Field",
		PkgName:       "schema",
		PkgImportName: "github.com/widmogrod/mkunion/x/schema",
		Fields: []*shape.FieldLike{
			{
				Name: "Name",
				Type: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
			},
			{
				Name: "Value",
				Type: &shape.RefName{
					Name:          "Schema",
					PkgName:       "schema",
					PkgImportName: "github.com/widmogrod/mkunion/x/schema",
				},
			},
		},
		Tags: map[string]shape.Tag{
			"serde": {
				Value: "json",
			},
		},
	}
}

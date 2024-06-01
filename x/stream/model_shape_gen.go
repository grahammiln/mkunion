// Code generated by mkunion. DO NOT EDIT.
package stream

import (
	"github.com/widmogrod/mkunion/x/shape"
)

func init() {
	shape.Register(EventTimeShape())
	shape.Register(FromBeginningShape())
	shape.Register(FromOffsetShape())
	shape.Register(ItemShape())
	shape.Register(OffsetShape())
	shape.Register(PullCMDShape())
	shape.Register(TopicShape())
}

//shape:shape

func PullCMDShape() shape.Shape {
	return &shape.UnionLike{
		Name:          "PullCMD",
		PkgName:       "stream",
		PkgImportName: "github.com/widmogrod/mkunion/x/stream",
		Variant: []shape.Shape{
			FromBeginningShape(),
			FromOffsetShape(),
		},
	}
}

func FromBeginningShape() shape.Shape {
	return &shape.StructLike{
		Name:          "FromBeginning",
		PkgName:       "stream",
		PkgImportName: "github.com/widmogrod/mkunion/x/stream",
		Fields: []*shape.FieldLike{
			{
				Name: "Topic",
				Type: &shape.RefName{
					Name:          "Topic",
					PkgName:       "stream",
					PkgImportName: "github.com/widmogrod/mkunion/x/stream",
				},
			},
		},
		Tags: map[string]shape.Tag{
			"mkunion": {
				Value: "PullCMD",
			},
		},
	}
}

func FromOffsetShape() shape.Shape {
	return &shape.StructLike{
		Name:          "FromOffset",
		PkgName:       "stream",
		PkgImportName: "github.com/widmogrod/mkunion/x/stream",
		Fields: []*shape.FieldLike{
			{
				Name: "Topic",
				Type: &shape.RefName{
					Name:          "Topic",
					PkgName:       "stream",
					PkgImportName: "github.com/widmogrod/mkunion/x/stream",
				},
			},
			{
				Name: "Offset",
				Type: &shape.PointerLike{
					Type: &shape.RefName{
						Name:          "Offset",
						PkgName:       "stream",
						PkgImportName: "github.com/widmogrod/mkunion/x/stream",
					},
				},
			},
		},
		Tags: map[string]shape.Tag{
			"mkunion": {
				Value: "PullCMD",
			},
		},
	}
}

//shape:shape
func TopicShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "Topic",
		PkgName:       "stream",
		PkgImportName: "github.com/widmogrod/mkunion/x/stream",
		IsAlias:       true,
		Type:          &shape.PrimitiveLike{Kind: &shape.StringLike{}},
	}
}

//shape:shape
func OffsetShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "Offset",
		PkgName:       "stream",
		PkgImportName: "github.com/widmogrod/mkunion/x/stream",
		Tags: map[string]shape.Tag{
			"serde": {
				Value: "json",
			},
		},
		Type: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
	}
}

//shape:shape
func EventTimeShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "EventTime",
		PkgName:       "stream",
		PkgImportName: "github.com/widmogrod/mkunion/x/stream",
		IsAlias:       true,
		Type: &shape.PrimitiveLike{
			Kind: &shape.NumberLike{
				Kind: &shape.Int64{},
			},
		},
	}
}

//shape:shape
func ItemShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Item",
		PkgName:       "stream",
		PkgImportName: "github.com/widmogrod/mkunion/x/stream",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
		},
		Fields: []*shape.FieldLike{
			{
				Name: "Topic",
				Type: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
			},
			{
				Name: "Key",
				Type: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
			},
			{
				Name: "Data",
				Type: &shape.RefName{
					Name:          "A",
					PkgName:       "",
					PkgImportName: "",
				},
			},
			{
				Name: "EventTime",
				Type: &shape.PointerLike{
					Type: &shape.RefName{
						Name:          "EventTime",
						PkgName:       "stream",
						PkgImportName: "github.com/widmogrod/mkunion/x/stream",
					},
				},
			},
			{
				Name: "Offset",
				Type: &shape.PointerLike{
					Type: &shape.RefName{
						Name:          "Offset",
						PkgName:       "stream",
						PkgImportName: "github.com/widmogrod/mkunion/x/stream",
					},
				},
			},
		},
	}
}

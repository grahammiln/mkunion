// Code generated by mkunion. DO NOT EDIT.
package workflow

import (
	"github.com/widmogrod/mkunion/x/shape"
)

func init() {
	shape.Register(FunctionInputShape())
	shape.Register(FunctionOutputShape())
}

//shape:shape
func FunctionInputShape() shape.Shape {
	return &shape.StructLike{
		Name:          "FunctionInput",
		PkgName:       "workflow",
		PkgImportName: "github.com/widmogrod/mkunion/x/workflow",
		Fields: []*shape.FieldLike{
			{
				Name: "Name",
				Type: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
			},
			{
				Name: "CallbackID",
				Type: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
			},
			{
				Name: "Args",
				Type: &shape.ListLike{
					Element: &shape.RefName{
						Name:          "Schema",
						PkgName:       "schema",
						PkgImportName: "github.com/widmogrod/mkunion/x/schema",
					},
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

//shape:shape
func FunctionOutputShape() shape.Shape {
	return &shape.StructLike{
		Name:          "FunctionOutput",
		PkgName:       "workflow",
		PkgImportName: "github.com/widmogrod/mkunion/x/workflow",
		Fields: []*shape.FieldLike{
			{
				Name: "Result",
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

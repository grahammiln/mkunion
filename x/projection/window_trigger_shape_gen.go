// Code generated by mkunion. DO NOT EDIT.
package projection

import (
	"github.com/widmogrod/mkunion/x/shape"
)

func init() {
	shape.Register(AtWatermarkShape())
	shape.Register(TriggerDescriptionShape())
}

//shape:shape

func TriggerDescriptionShape() shape.Shape {
	return &shape.UnionLike{
		Name:          "TriggerDescription",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		Variant: []shape.Shape{
			AtWatermarkShape(),
		},
	}
}

func AtWatermarkShape() shape.Shape {
	return &shape.StructLike{
		Name:          "AtWatermark",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		Fields: []*shape.FieldLike{
			{
				Name: "Timestamp",
				Type: &shape.PrimitiveLike{
					Kind: &shape.NumberLike{
						Kind: &shape.Int64{},
					},
				},
			},
		},
	}
}

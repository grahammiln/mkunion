// Code generated by mkunion. DO NOT EDIT.
package f

import (
	"github.com/widmogrod/mkunion/x/shape"
)

func init() {
	shape.Register(EitherShape())
	shape.Register(ErrShape())
	shape.Register(LeftShape())
	shape.Register(NoneShape())
	shape.Register(OkShape())
	shape.Register(OptionShape())
	shape.Register(ResultShape())
	shape.Register(RightShape())
	shape.Register(SomeShape())
}

//shape:shape

func EitherShape() shape.Shape {
	return &shape.UnionLike{
		Name:          "Either",
		PkgName:       "f",
		PkgImportName: "github.com/widmogrod/mkunion/f",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
			shape.TypeParam{
				Name: "B",
				Type: &shape.Any{},
			},
		},
		Variant: []shape.Shape{
			LeftShape(),
			RightShape(),
		},
	}
}

func LeftShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Left",
		PkgName:       "f",
		PkgImportName: "github.com/widmogrod/mkunion/f",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
			shape.TypeParam{
				Name: "B",
				Type: &shape.Any{},
			},
		},
		Fields: []*shape.FieldLike{
			{
				Name: "Value",
				Type: &shape.RefName{
					Name:          "A",
					PkgName:       "",
					PkgImportName: "",
				},
			},
		},
		Tags: map[string]shape.Tag{
			"mkunion_union_name": {
				Value: "Either",
			},
		},
	}
}

func RightShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Right",
		PkgName:       "f",
		PkgImportName: "github.com/widmogrod/mkunion/f",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
			shape.TypeParam{
				Name: "B",
				Type: &shape.Any{},
			},
		},
		Fields: []*shape.FieldLike{
			{
				Name: "Value",
				Type: &shape.RefName{
					Name:          "B",
					PkgName:       "",
					PkgImportName: "",
				},
			},
		},
		Tags: map[string]shape.Tag{
			"mkunion_union_name": {
				Value: "Either",
			},
		},
	}
}

//shape:shape

func OptionShape() shape.Shape {
	return &shape.UnionLike{
		Name:          "Option",
		PkgName:       "f",
		PkgImportName: "github.com/widmogrod/mkunion/f",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
		},
		Variant: []shape.Shape{
			SomeShape(),
			NoneShape(),
		},
	}
}

func SomeShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Some",
		PkgName:       "f",
		PkgImportName: "github.com/widmogrod/mkunion/f",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
		},
		Fields: []*shape.FieldLike{
			{
				Name: "Value",
				Type: &shape.RefName{
					Name:          "A",
					PkgName:       "",
					PkgImportName: "",
				},
			},
		},
		Tags: map[string]shape.Tag{
			"mkunion_union_name": {
				Value: "Option",
			},
		},
	}
}

func NoneShape() shape.Shape {
	return &shape.StructLike{
		Name:          "None",
		PkgName:       "f",
		PkgImportName: "github.com/widmogrod/mkunion/f",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
		},
		Tags: map[string]shape.Tag{
			"mkunion_union_name": {
				Value: "Option",
			},
		},
	}
}

//shape:shape

func ResultShape() shape.Shape {
	return &shape.UnionLike{
		Name:          "Result",
		PkgName:       "f",
		PkgImportName: "github.com/widmogrod/mkunion/f",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
			shape.TypeParam{
				Name: "E",
				Type: &shape.Any{},
			},
		},
		Variant: []shape.Shape{
			OkShape(),
			ErrShape(),
		},
	}
}

func OkShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Ok",
		PkgName:       "f",
		PkgImportName: "github.com/widmogrod/mkunion/f",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
			shape.TypeParam{
				Name: "E",
				Type: &shape.Any{},
			},
		},
		Fields: []*shape.FieldLike{
			{
				Name: "Value",
				Type: &shape.RefName{
					Name:          "A",
					PkgName:       "",
					PkgImportName: "",
				},
			},
		},
		Tags: map[string]shape.Tag{
			"mkunion_union_name": {
				Value: "Result",
			},
		},
	}
}

func ErrShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Err",
		PkgName:       "f",
		PkgImportName: "github.com/widmogrod/mkunion/f",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
			shape.TypeParam{
				Name: "E",
				Type: &shape.Any{},
			},
		},
		Fields: []*shape.FieldLike{
			{
				Name: "Value",
				Type: &shape.RefName{
					Name:          "E",
					PkgName:       "",
					PkgImportName: "",
				},
			},
		},
		Tags: map[string]shape.Tag{
			"mkunion_union_name": {
				Value: "Result",
			},
		},
	}
}

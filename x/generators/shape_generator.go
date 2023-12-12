package generators

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/widmogrod/mkunion/x/shape"
	"strings"
	"text/template"
)

var (
	//go:embed shape_generator.go.tmpl
	shapeTmpl string
)

func NewShapeGenerator(union shape.UnionLike, helper *Helpers) *ShapeGenerator {
	return &ShapeGenerator{
		Union:    union,
		template: template.Must(template.New("shape_generator.go.tmpl").Funcs(helper.Func()).Parse(shapeTmpl)),
	}
}

type ShapeGenerator struct {
	Union    shape.UnionLike
	template *template.Template
}

func (g *ShapeGenerator) ident(d int) string {
	return strings.Repeat("\t", d)
}

func (g *ShapeGenerator) padLeft(d int, s string) string {
	// pad each new line with \t
	return strings.ReplaceAll(s, "\n", "\n"+g.ident(d))
}

func (g *ShapeGenerator) VariantName(x shape.Shape) string {
	return TemplateHelperShapeVariantToName(x)
}

func TemplateHelperShapeVariantToName(x shape.Shape) string {
	return shape.MustMatchShape(
		x,
		func(x *shape.Any) string {
			panic(fmt.Errorf("generators.TemplateHelperShapeVariantToName: %T not suported", x))
		},
		func(x *shape.RefName) string {
			return x.Name
		},
		func(x *shape.BooleanLike) string {
			if shape.IsNamed(x) {
				return x.Named.Name
			}
			panic(fmt.Errorf("generators.TemplateHelperShapeVariantToName: expects only named shape: %#v", x))
		},
		func(x *shape.StringLike) string {
			if shape.IsNamed(x) {
				return x.Named.Name
			}
			panic(fmt.Errorf("generators.TemplateHelperShapeVariantToName: expects only named shape: %#v", x))
		},
		func(x *shape.NumberLike) string {
			if shape.IsNamed(x) {
				return x.Named.Name
			}
			panic(fmt.Errorf("generators.TemplateHelperShapeVariantToName: expects only named shape: %#v", x))
		},
		func(x *shape.ListLike) string {
			if shape.IsNamed(x) {
				return x.Named.Name
			}
			panic(fmt.Errorf("generators.TemplateHelperShapeVariantToName: expects only named shape: %#v", x))
		},
		func(x *shape.MapLike) string {
			if shape.IsNamed(x) {
				return x.Named.Name
			}
			panic(fmt.Errorf("generators.TemplateHelperShapeVariantToName: expects only named shape: %#v", x))
		},
		func(x *shape.StructLike) string {
			return x.Name
		},
		func(x *shape.UnionLike) string {
			return x.Name
		},
	)
}

func (g *ShapeGenerator) ShapeToString(x shape.Shape, depth int) string {
	return shape.MustMatchShape(
		x,
		func(x *shape.Any) string {
			return g.padLeft(depth, `&shape.Any{}`)
		},
		func(x *shape.RefName) string {
			result := &bytes.Buffer{}

			fmt.Fprintf(result, "&shape.RefName{\n")
			fmt.Fprintf(result, "\tName: %q,\n", x.Name)
			fmt.Fprintf(result, "\tPkgName: %q,\n", x.PkgName)
			fmt.Fprintf(result, "\tPkgImportName: %q,\n", x.PkgImportName)
			fmt.Fprintf(result, "}")

			return g.padLeft(depth, result.String())
		},
		func(x *shape.BooleanLike) string {
			result := &bytes.Buffer{}
			if shape.IsNamed(x) {
				fmt.Fprintf(result, "&shape.BooleanLike{\n")
				g.fprintNamedFields(result, x.Named, 1)
				fmt.Fprintf(result, "}")
			} else {
				fmt.Fprintf(result, "&shape.BooleanLike{}")
			}

			return g.padLeft(depth, result.String())
		},
		func(x *shape.StringLike) string {
			result := &bytes.Buffer{}
			if shape.IsNamed(x) {
				fmt.Fprintf(result, "&shape.StringLike{\n")
				g.fprintNamedFields(result, x.Named, 1)
				fmt.Fprintf(result, "}")
			} else {
				fmt.Fprintf(result, "&shape.StringLike{}")
			}

			return g.padLeft(depth, result.String())
		},
		func(x *shape.NumberLike) string {
			result := &bytes.Buffer{}
			if shape.IsNamed(x) {
				fmt.Fprintf(result, "&shape.NumberLike{\n")
				g.fprintNamedFields(result, x.Named, 1)
				fmt.Fprintf(result, "}")
			} else {
				fmt.Fprintf(result, "&shape.NumberLike{}")
			}

			return g.padLeft(depth, result.String())
		},
		func(x *shape.ListLike) string {
			result := &bytes.Buffer{}

			fmt.Fprintf(result, "&shape.ListLike{\n")
			fmt.Fprintf(result, "\tElement: %s,\n", strings.TrimLeft(g.ShapeToString(x.Element, 1), "\t"))
			fmt.Fprintf(result, "\tElementIsPointer: %v,\n", x.ElementIsPointer)

			if shape.IsNamed(x) {
				g.fprintNamedFields(result, x.Named, 2)
			}

			fmt.Fprintf(result, "}")

			return g.padLeft(depth, result.String())
		},
		func(x *shape.MapLike) string {
			result := &bytes.Buffer{}

			fmt.Fprintf(result, "&shape.MapLike{\n")
			fmt.Fprintf(result, "\tKey: %s,\n", strings.TrimLeft(g.ShapeToString(x.Key, 1), "\t"))
			fmt.Fprintf(result, "\tKeyIsPointer: %v,\n", x.KeyIsPointer)
			fmt.Fprintf(result, "\tVal: %s,\n", strings.TrimLeft(g.ShapeToString(x.Val, 1), "\t"))
			fmt.Fprintf(result, "\tValIsPointer: %v,\n", x.ValIsPointer)

			if shape.IsNamed(x) {
				g.fprintNamedFields(result, x.Named, 2)
			}

			fmt.Fprintf(result, "}")

			return g.padLeft(depth, result.String())
		},
		func(x *shape.StructLike) string {
			result := &bytes.Buffer{}

			fmt.Fprintf(result, "&shape.StructLike{\n")
			fmt.Fprintf(result, "\tName: %q,\n", x.Name)
			fmt.Fprintf(result, "\tPkgName: %q,\n", x.PkgName)
			fmt.Fprintf(result, "\tPkgImportName: %q,\n", x.PkgImportName)

			if len(x.Fields) > 0 {
				fmt.Fprintf(result, "\tFields: []*shape.FieldLike{\n")
				for _, field := range x.Fields {
					fmt.Fprintf(result, "\t\t{\n")
					fmt.Fprintf(result, "\t\t\tName: %q,\n", field.Name)
					fmt.Fprintf(result, "\t\t\tType: %s,\n", strings.TrimLeft(g.ShapeToString(field.Type, 3), "\t"))
					fmt.Fprintf(result, "\t\t},\n")
				}
				fmt.Fprintf(result, "\t},\n")
			}
			fmt.Fprintf(result, "}")

			return g.padLeft(depth, result.String())
		},
		func(x *shape.UnionLike) string {
			result := &bytes.Buffer{}

			fmt.Fprintf(result, "&shape.UnionLike{\n")
			fmt.Fprintf(result, "\tName: %q,\n", x.Name)
			fmt.Fprintf(result, "\tPkgName: %q,\n", x.PkgName)
			fmt.Fprintf(result, "\tPkgImportName: %q,\n", x.PkgImportName)

			if len(x.Variant) > 0 {
				fmt.Fprintf(result, "\tVariant: []shape.Shape{\n")
				for _, variant := range x.Variant {
					fmt.Fprintf(result, "\t\t%s,\n", strings.TrimLeft(g.ShapeToString(variant, 2), "\t"))
				}
				fmt.Fprintf(result, "\t},\n")
			}

			fmt.Fprintf(result, "}")

			return g.padLeft(depth, result.String())
		},
	)
}

func (g *ShapeGenerator) fprintNamedFields(result *bytes.Buffer, x *shape.Named, depth int) {
	fmt.Fprintf(result, strings.Repeat("\t", depth)+"Named: &shape.Named{\n")
	fmt.Fprintf(result, strings.Repeat("\t", depth)+"\tName: %q,\n", x.Name)
	fmt.Fprintf(result, strings.Repeat("\t", depth)+"\tPkgName: %q,\n", x.PkgName)
	fmt.Fprintf(result, strings.Repeat("\t", depth)+"\tPkgImportName: %q,\n", x.PkgImportName)
	fmt.Fprintf(result, strings.Repeat("\t", depth)+"},\n")
}

func (g *ShapeGenerator) Generate() ([]byte, error) {
	result := &bytes.Buffer{}
	err := g.template.ExecuteTemplate(result, "shape_generator.go.tmpl", g)
	if err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}
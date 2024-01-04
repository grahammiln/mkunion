package generators

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/widmogrod/mkunion/x/shape"
	"text/template"
)

var (
	//go:embed serde_json_union.go.tmpl
	deserJSONTmpl string
)

func NewSerdeJSONUnion(union *shape.UnionLike) *SerdeJSONUnion {
	return &SerdeJSONUnion{
		Union:                 union,
		template:              template.Must(template.New("serde_json_union.go.tmpl").Parse(deserJSONTmpl)),
		skipImportsAndPackage: false,
		pkgUsed: PkgMap{
			"json":   "encoding/json",
			"fmt":    "fmt",
			"shared": "github.com/widmogrod/mkunion/x/shared",
		},
	}
}

type SerdeJSONUnion struct {
	Union                 *shape.UnionLike
	template              *template.Template
	skipImportsAndPackage bool
	pkgUsed               PkgMap
}

func (g *SerdeJSONUnion) SkipImportsAndPackage(x bool) {
	g.skipImportsAndPackage = x
}

func (g *SerdeJSONUnion) GenerateImports(pkgMap PkgMap) (string, error) {
	return GenerateImports(pkgMap), nil
}

func (g *SerdeJSONUnion) ExtractImports(x shape.Shape) PkgMap {
	pkgMap := shape.ExtractPkgImportNames(x)
	if pkgMap == nil {
		pkgMap = make(map[string]string)
	}

	// add default and necessary imports
	pkgMap = MergePkgMaps(pkgMap, g.pkgUsed)

	// remove self from importing
	delete(pkgMap, shape.ToGoPkgName(x))
	return pkgMap
}

func (g *SerdeJSONUnion) IsStruct(x shape.Shape) bool {
	_, ok := x.(*shape.StructLike)
	return ok
}

func (g *SerdeJSONUnion) VariantName(x shape.Shape) string {
	return TemplateHelperShapeVariantToName(x)
}

func (g *SerdeJSONUnion) JSONVariantName(x shape.Shape) string {
	return shape.MatchShapeR1(
		x,
		func(y *shape.Any) string {
			panic(fmt.Errorf("generators.JSONVariantName: %T not suported", y))
		},
		func(y *shape.RefName) string {
			return fmt.Sprintf("%s.%s", y.PkgName, y.Name)
		},
		func(x *shape.PointerLike) string {
			return g.JSONVariantName(x.Type)
		},
		func(y *shape.AliasLike) string {
			return fmt.Sprintf("%s.%s", y.PkgName, y.Name)
		},
		func(y *shape.PrimitiveLike) string {
			panic(fmt.Errorf("generators.JSONVariantName: must be named %T", y))
		},
		func(y *shape.ListLike) string {
			panic(fmt.Errorf("generators.JSONVariantName: must be named %T", y))
		},
		func(y *shape.MapLike) string {
			panic(fmt.Errorf("generators.JSONVariantName: must be named %T", y))
		},
		func(y *shape.StructLike) string {
			return fmt.Sprintf("%s.%s", y.PkgName, y.Name)
		},
		func(y *shape.UnionLike) string {
			return fmt.Sprintf("%s.%s", y.PkgName, y.Name)
		},
	)
}

func (g *SerdeJSONUnion) Generate() ([]byte, error) {
	body := &bytes.Buffer{}
	err := g.template.ExecuteTemplate(body, "serde_json_union.go.tmpl", g)
	if err != nil {
		return nil, err
	}

	head := &bytes.Buffer{}
	if !g.skipImportsAndPackage {
		head.WriteString(fmt.Sprintf("package %s\n\n", shape.ToGoPkgName(g.Union)))

		pkgMap := g.ExtractImports(g.Union)
		impPart, err := g.GenerateImports(pkgMap)
		if err != nil {
			return nil, fmt.Errorf("generators.SerdeJSONTagged.Generate: when generating imports %w", err)
		}
		head.WriteString(impPart)
	}

	if head.Len() > 0 {
		head.WriteString(body.String())
		return head.Bytes(), nil
	} else {
		return body.Bytes(), nil
	}
}

func (g *SerdeJSONUnion) Serde(x shape.Shape) string {
	serde := NewSerdeJSONTagged(x)
	serde.SkipImportsAndPackage(true)
	result, err := serde.Generate()
	if err != nil {
		panic(err)
	}

	g.pkgUsed = MergePkgMaps(g.pkgUsed, serde.ExtractImports(x))

	return result
}

func (g *SerdeJSONUnion) MatchFuncName(x *shape.UnionLike, returns int) string {
	return MatchUnionFuncName(x, returns)
}

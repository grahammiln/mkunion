package generators

import (
	"fmt"
	"github.com/widmogrod/mkunion/x/shape"
	"github.com/widmogrod/mkunion/x/shared"
	"sort"
	"strings"
)

const (
	Header = `// Code generated by ` + shared.Program + `. DO NOT EDIT.`
)

type Generator interface {
	Generate() ([]byte, error)
}

type PkgMap = map[string]string
type InitFuncs = []string

func GenerateImports(pkgMap PkgMap) string {
	if len(pkgMap) == 0 {
		return ""
	}

	result := &strings.Builder{}
	result.WriteString("import (\n")

	var sortedImportNames []string
	for _, pkgImportName := range pkgMap {
		sortedImportNames = append(sortedImportNames, pkgImportName)
	}
	sort.Strings(sortedImportNames)

	for _, pkgImportName := range sortedImportNames {
		result.WriteString(fmt.Sprintf("\t\"%s\"\n", pkgImportName))
	}
	result.WriteString(")\n\n")

	return result.String()
}

func GenerateInitFunc(inits InitFuncs) string {
	if len(inits) == 0 {
		return ""
	}

	result := &strings.Builder{}
	result.WriteString("func init() {\n")

	for _, init := range inits {
		result.WriteString(fmt.Sprintf("\t%s\n", init))
	}
	result.WriteString("}\n\n")

	return result.String()
}

func MergePkgMaps(maps ...PkgMap) PkgMap {
	result := make(PkgMap)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func padLeftTabs(n int, s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if line == "" {
			// don't add tabs to empty lines
			continue
		}
		lines[i] = strings.Repeat("\t", n) + line
	}
	return strings.Join(lines, "\n")
}

func padLeftTabs2(n int, s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if i == 0 {
			continue
		}
		if line == "" {
			// don't add tabs to empty lines
			continue
		}
		lines[i] = strings.Repeat("\t", n) + line
	}
	return strings.Join(lines, "\n")
}

func TypeNameIfSupports(s shape.Shape) (string, bool) {
	return shape.MatchShapeR2(
		s,
		func(x *shape.Any) (string, bool) {
			return "", false
		},
		func(x *shape.RefName) (string, bool) {
			return "", false
		},
		func(x *shape.PointerLike) (string, bool) {
			return TypeNameIfSupports(x.Type)
		},
		func(x *shape.AliasLike) (string, bool) {
			return x.Name, true
		},
		func(x *shape.PrimitiveLike) (string, bool) {
			return "", false
		},
		func(x *shape.ListLike) (string, bool) {
			return "", false
		},
		func(x *shape.MapLike) (string, bool) {
			return "", false
		},
		func(x *shape.StructLike) (string, bool) {
			return x.Name, true
		},
		func(x *shape.UnionLike) (string, bool) {
			return x.Name, true
		},
	)
}

func TemplateHelperShapeVariantToName(x shape.Shape) string {
	return shape.MatchShapeR1(
		x,
		func(x *shape.Any) string {
			panic(fmt.Errorf("generators.TemplateHelperShapeVariantToName: %T not suported", x))
		},
		func(x *shape.RefName) string {
			return x.Name
		},
		func(x *shape.PointerLike) string {
			return TemplateHelperShapeVariantToName(x.Type)
		},
		func(x *shape.AliasLike) string {
			return x.Name
		},
		func(x *shape.PrimitiveLike) string {
			panic(fmt.Errorf("generators.TemplateHelperShapeVariantToName: expects only named shape: %#v", x))
		},
		func(x *shape.ListLike) string {
			panic(fmt.Errorf("generators.TemplateHelperShapeVariantToName: expects only named shape: %#v", x))
		},
		func(x *shape.MapLike) string {
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

package mkunion

import (
	"bytes"
	_ "embed"
	"fmt"
	log "github.com/sirupsen/logrus"
	"text/template"
)

const (
	Program = "mkunion"
	Header  = `// Code generated by ` + Program + `. DO NOT EDIT.`
)

type generatorOptions struct {
	Header      string
	PackageName string

	bufferImports bool
}

type GenerateOption func(o *generatorOptions)

func WithBufferedImports() GenerateOption {
	return func(o *generatorOptions) {
		o.bufferImports = true
	}
}

func WithPackageName(name string) GenerateOption {
	return func(o *generatorOptions) {
		o.PackageName = name
	}
}

type Generator interface {
	Generate() ([]byte, error)
}

func NewHelper(fs ...GenerateOption) *Helpers {
	options := generatorOptions{
		Header:      Header,
		PackageName: "main",
	}

	for _, f := range fs {
		f(&options)
	}

	return &Helpers{
		options: options,
		imports: nil,
	}
}

type Helpers struct {
	options generatorOptions

	imports []string
}

func (help *Helpers) RenderHeader() string {
	if help.options.bufferImports {
		return ""
	}

	return help.RenderBufferedHeader()
}

func (help *Helpers) RenderBufferedHeader() string {
	result := &bytes.Buffer{}
	fmt.Fprintf(result, "%s\n", help.options.Header)
	fmt.Fprintf(result, "package %s\n\n", help.options.PackageName)
	return result.String()
}

func (help *Helpers) RenderBufferedImport() string {
	result := &bytes.Buffer{}
	log.Debugf("render import %v", help.imports)
	for _, name := range help.imports {
		fmt.Fprintf(result, "import %q\n", name)
	}

	return result.String()
}

func (help *Helpers) RenderImport(importName ...string) string {
	if len(importName) == 0 {
		return ""
	}

	if help.options.bufferImports {
		log.Debugf("buffer import %v", importName)
		help.imports = append(help.imports, importName...)
		return ""
	}

	result := &bytes.Buffer{}
	for _, name := range importName {
		fmt.Fprintf(result, "import %q\n", name)
	}

	result.WriteString("\n")

	return result.String()
}

func (help *Helpers) Func() template.FuncMap {
	return map[string]any{
		"RenderHeader": help.RenderHeader,
		"RenderImport": help.RenderImport,
		"GenIntSlice": func(from, to int) []int {
			var result []int
			for i := from; i <= to; i++ {
				result = append(result, i)
			}
			return result
		},
		"Add": func(a, b int) int {
			return a + b
		},
	}
}

var (
	//go:embed visitor_generator.go.tmpl
	visitorTmpl string
)

func NewVisitorGenerator(name string, types []string, helper *Helpers) *VisitorGenerator {
	return &VisitorGenerator{
		Name:     name,
		Types:    types,
		Helper:   helper,
		template: template.Must(template.New("main").Funcs(helper.Func()).Parse(visitorTmpl)),
	}
}

type VisitorGenerator struct {
	Types    []string
	Name     string
	Helper   *Helpers
	template *template.Template
}

func (g *VisitorGenerator) Generate() ([]byte, error) {
	result := &bytes.Buffer{}

	err := g.template.ExecuteTemplate(result, "main", g)

	if err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}

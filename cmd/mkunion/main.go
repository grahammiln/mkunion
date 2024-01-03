package main

import (
	"bytes"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/widmogrod/mkunion/x/generators"
	"github.com/widmogrod/mkunion/x/shape"
	"github.com/widmogrod/mkunion/x/shared"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
)

type extractImports interface {
	ExtractImports(x shape.Shape) generators.PkgMap
	SkipImportsAndPackage(x bool)
}

type extractInitFuncs interface {
	ExtractImportFuncs(s shape.Shape) []string
	SkipInitFunc(flag bool)
}

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	// set log level to error
	log.SetOutput(os.Stderr)
	log.SetLevel(log.ErrorLevel)

	var app *cli.App
	app = &cli.App{
		Name:                   shared.Program,
		Description:            "VisitorGenerator union type and visitor pattern gor golang",
		EnableBashCompletion:   true,
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Required: false,
			},
			&cli.StringFlag{
				Name:     "skip-extension",
				Aliases:  []string{"skip-ext"},
				Value:    "",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "include-extension",
				Aliases:  []string{"inc-ext"},
				Required: false,
			},
			&cli.StringSliceFlag{
				Name:      "input-go-file",
				Aliases:   []string{"i", "input"},
				Usage:     `When not provided, it will try to use GOFILE environment variable, used when combined with //go:generate mkunion -name=MyUnionType`,
				TakesFile: true,
			},
			&cli.BoolFlag{
				Name:     "verbose",
				Aliases:  []string{"v"},
				Required: false,
				Value:    false,
			},
			&cli.BoolFlag{
				Name:     "no-compact",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("verbose") {
				log.SetLevel(log.DebugLevel)
			}

			sourcePaths := c.StringSlice("input-go-file")
			if len(sourcePaths) == 0 && os.Getenv("GOFILE") != "" {
				cwd, _ := syscall.Getwd()
				sourceName := path.Base(os.Getenv("GOFILE"))
				sourcePaths = []string{
					path.Join(cwd, sourceName),
				}
			}

			if len(sourcePaths) == 0 {
				// show usage
				cli.ShowAppHelpAndExit(c, 1)
			}

			for _, sourcePath := range sourcePaths {
				cwd := path.Dir(sourcePath)
				sourceName := path.Base(sourcePath)
				baseName := strings.TrimSuffix(sourceName, path.Ext(sourceName))

				// file name without extension
				inferred, err := shape.InferFromFile(sourcePath)
				if err != nil {
					return err
				}

				var unions []*shape.UnionLike
				for _, unionName := range c.StringSlice("name") {
					union := inferred.RetrieveUnion(unionName)
					if union == nil {
						return fmt.Errorf("union %s not found in %s", unionName, sourcePath)
					}

					unions = append(unions, union)
				}

				if len(unions) == 0 {
					unions = inferred.RetrieveUnions()
				}

				for _, union := range unions {
					if union == nil {
						return fmt.Errorf("union %s not found in %s", union.Name, sourcePath)
					}

					jsonGenerator := generators.SerdeJSONUnion(union)
					shapeGenerator := generators.NewShapeUnion(union)
					visitor := generators.NewVisitorGenerator(union)

					// ensures that order of generators2 is always the same
					generatorsList := []string{
						"visitor",
						"shape",
						"json",
					}

					generators2 := map[string]generators.Generator{
						"visitor": visitor,
						"shape":   shapeGenerator,
						"json":    jsonGenerator,
					}

					skipExtension := strings.Split(c.String("skip-extension"), ",")
					includeExtension := strings.Split(c.String("include-extension"), ",")
					if len(includeExtension) > 0 {
						for _, includeName := range includeExtension {
							for i, skipName := range skipExtension {
								if skipName == includeName {
									log.Infof("include extension, that was skipName %s", includeName)
									skipExtension = append(skipExtension[:i], skipExtension[i+1:]...)
								}
							}
						}
					}

					for _, name := range skipExtension {
						log.Infof("skip extension %s", name)
						delete(generators2, name)
					}

					if c.Bool("no-compact") {
						for _, name := range generatorsList {
							g, ok := generators2[name]
							if !ok {
								continue
							}

							b, err := g.Generate()
							if err != nil {
								return fmt.Errorf("failed to generate %s for %s in %s: %w", name, union.Name, sourcePath, err)
							}

							fileName := baseName + "_" + shared.Program + "_" + strings.ToLower(union.Name) + "_" + name + ".go"
							log.Infof("writing %s", fileName)

							err = os.WriteFile(path.Join(cwd, fileName), b, 0644)
							if err != nil {
								return fmt.Errorf("failed to write %s for %s in %s: %w", name, union.Name, sourcePath, err)
							}
						}
					} else {
						packageName := union.PkgName
						pkgMap := make(generators.PkgMap)
						initFunc := make(generators.InitFuncs, 0, 0)
						shapesContents := bytes.Buffer{}

						for _, name := range generatorsList {
							g, ok := generators2[name]
							if !ok {
								continue
							}
							if gen, ok := g.(extractImports); ok {
								gen.SkipImportsAndPackage(true)
							}
							if gen, ok := g.(extractInitFuncs); ok {
								gen.SkipInitFunc(true)
							}

							b, err := g.Generate()
							if err != nil {
								return fmt.Errorf("failed to generate %s for %s in %s: %w", name, union.Name, sourcePath, err)
							}

							if gen, ok := g.(extractImports); ok {
								gen.ExtractImports(union)
								pkgMap = generators.MergePkgMaps(pkgMap,
									gen.ExtractImports(union),
								)
							}
							if gen, ok := g.(extractInitFuncs); ok {
								initFunc = append(initFunc, gen.ExtractImportFuncs(union)...)
							}

							shapesContents.WriteString(fmt.Sprintf("//mkunion-extension:%s\n", name))
							shapesContents.Write(b)
							shapesContents.WriteString("\n")
						}

						contents := "// Code generated by mkunion. DO NOT EDIT.\n"
						contents += fmt.Sprintf("package %s\n\n", packageName)
						contents += generators.GenerateImports(pkgMap)
						contents += generators.GenerateInitFunc(initFunc)
						contents += shapesContents.String()

						fileName := path.Join(
							path.Dir(sourcePath),
							fmt.Sprintf("%s_%s_gen.go", baseName, strings.ToLower(union.Name)),
						)

						log.Infof("writing %s", fileName)
						err = os.WriteFile(fileName, []byte(contents), 0644)
						if err != nil {
							return fmt.Errorf("failed to write(2) %s for %s in %s: %w", "gen", union.Name, sourcePath, err)
						}
					}
				}
			}

			return nil
		},
		Commands: []*cli.Command{
			{
				Name: "match",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					cwd, _ := syscall.Getwd()
					sourceName := path.Base(os.Getenv("GOFILE"))
					sourcePath := path.Join(cwd, sourceName)

					baseName := strings.TrimSuffix(sourceName, path.Ext(sourceName))

					// file name without extension
					inferred, err := generators.InferDeriveFuncMatchFromFile(sourcePath)
					if err != nil {
						return err
					}

					specName := c.String("name")
					spec, err := inferred.MatchSpec(specName)
					if err != nil {
						return err
					}

					derived := generators.DeriveFuncMatchGenerator{
						Header:      "// Code generated by mkunion. DO NOT EDIT.",
						PackageName: inferred.PackageName,
						MatchSpec:   *spec,
					}

					b, err := derived.Generate()
					if err != nil {
						return err
					}
					err = os.WriteFile(path.Join(
						cwd,
						baseName+"_match_"+strings.ToLower(derived.MatchSpec.Name)+".go"), b, 0644)
					if err != nil {
						return fmt.Errorf("failed to write %s for %s in %s: %w", "gen", derived.MatchSpec.Name, sourcePath, err)
					}

					return nil
				},
			},
			{
				Name: "serde",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "type",
						DefaultText: "json",
					},
					&cli.StringSliceFlag{
						Name:      "input-go-file",
						Aliases:   []string{"i", "input"},
						Usage:     `When not provided, it will try to use GOFILE environment variable, used when combined with //go:generate mkunion -name=MyUnionType`,
						TakesFile: true,
					},
					&cli.BoolFlag{
						Name:     "verbose",
						Aliases:  []string{"v"},
						Required: false,
						Value:    false,
					},
				},
				Action: func(c *cli.Context) error {
					if c.Bool("verbose") {
						log.SetLevel(log.DebugLevel)
					}

					sourcePaths := c.StringSlice("input-go-file")
					if len(sourcePaths) == 0 && os.Getenv("GOFILE") != "" {
						cwd, _ := syscall.Getwd()
						sourceName := path.Base(os.Getenv("GOFILE"))
						sourcePaths = []string{
							path.Join(cwd, sourceName),
						}
					}

					if len(sourcePaths) == 0 {
						// show usage
						cli.ShowAppHelpAndExit(c, 1)
					}

					for _, sourcePath := range sourcePaths {
						inferred, err := shape.InferFromFile(sourcePath)
						if err != nil {
							return fmt.Errorf("failed inferring shape in %s; %w", sourcePath, err)
						}

						shapes := inferred.RetrieveShapesTaggedAs("serde")
						if len(shapes) == 0 {
							log.Infof("no shape found in %s", sourcePath)
							continue
						}

						packageName := "main"
						pkgMap := make(generators.PkgMap)
						initFunc := make(generators.InitFuncs, 0, 0)
						shapesContents := bytes.Buffer{}

						for _, x := range shapes {
							packageName = shape.ToGoPkgName(x)
							genSerde := generators.NewSerdeJSONTagged(x)
							genSerde.SkipImportsAndPackage(true)

							genShape := generators.NewShapeTagged(x)
							genShape.SkipImportsAndPackage(true)
							genShape.SkipInitFunc(true)

							contents := "//serde:json\n"
							contents, err = genSerde.Generate()
							if err != nil {
								return fmt.Errorf("failed to generate json serde for %s in %s: %w", shape.ToGoTypeName(x), sourcePath, err)
							}
							shapesContents.WriteString(contents)

							contents = "//shape\n"
							contents, err = genShape.Generate()
							if err != nil {
								return fmt.Errorf("failed to generate shape for %s in %s: %w", shape.ToGoTypeName(x), sourcePath, err)
							}
							shapesContents.WriteString(contents)

							pkgMap = generators.MergePkgMaps(pkgMap,
								genSerde.ExtractImports(x),
								genShape.ExtractImports(x),
							)

							initFunc = append(initFunc, genShape.ExtractImportFuncs(x)...)
						}

						contents := "// Code generated by mkunion. DO NOT EDIT.\n"
						contents += fmt.Sprintf("package %s\n\n", packageName)
						contents += generators.GenerateImports(pkgMap)
						contents += generators.GenerateInitFunc(initFunc)
						contents += shapesContents.String()

						sourceName := path.Base(sourcePath)
						baseName := strings.TrimSuffix(sourceName, path.Ext(sourceName))
						fileName := path.Join(
							path.Dir(sourcePath),
							fmt.Sprintf("%s_serde_gen.go", baseName),
						)

						log.Infof("writing %s", fileName)
						err = os.WriteFile(fileName, []byte(contents), 0644)
						if err != nil {
							return fmt.Errorf("failed to write serialiser in %s: %w", sourcePath, err)
						}

						return nil
					}

					return nil
				},
			},
			{
				Name: "shape-export",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "language",
						Aliases:     []string{"lang"},
						DefaultText: "typescript",
					},
					&cli.StringFlag{
						Name:    "output-dir",
						Aliases: []string{"o", "output"},
					},
					&cli.StringSliceFlag{
						Name:      "input-go-file",
						Aliases:   []string{"i", "input"},
						Usage:     `When not provided, it will try to use GOFILE environment variable, used when combined with //go:generate mkunion -name=MyUnionType`,
						TakesFile: true,
					},
					&cli.BoolFlag{
						Name:     "verbose",
						Aliases:  []string{"v"},
						Required: false,
						Value:    false,
					},
				},
				Action: func(c *cli.Context) error {
					if c.Bool("verbose") {
						log.SetLevel(log.DebugLevel)
					}

					sourcePaths := c.StringSlice("input-go-file")
					if len(sourcePaths) == 0 && os.Getenv("GOFILE") != "" {
						cwd, _ := syscall.Getwd()
						sourceName := path.Base(os.Getenv("GOFILE"))
						sourcePaths = []string{
							path.Join(cwd, sourceName),
						}
					}

					if len(sourcePaths) == 0 {
						// show usage
						cli.ShowAppHelpAndExit(c, 1)
					}

					tsr := shape.NewTypeScriptRenderer()
					for _, sourcePath := range sourcePaths {
						// file name without extension
						inferred, err := shape.InferFromFile(sourcePath)
						if err != nil {
							return err
						}

						for _, x := range inferred.RetrieveShapes() {
							tsr.AddShape(x)
							tsr.FollowRef(x)
						}
					}

					tsr.FollowImports()

					err := tsr.WriteToDir(c.String("output-dir"))
					if err != nil {
						return fmt.Errorf("failed to write to dir %s: %w", c.String("output-dir"), err)
					}

					return nil
				},
			},
		},
	}

	err := app.RunContext(ctx, os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

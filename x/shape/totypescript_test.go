package shape

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTypeScriptSchemaGeneration(t *testing.T) {
	inferred, err := InferFromFile("testasset/type_example.go")
	if err != nil {
		t.Fatal(err)
	}

	union := inferred.RetrieveUnion("Example")

	tsr := NewTypeScriptRenderer()
	tsr.AddUnion(&union)

	tsr.AddStruct(&StructLike{
		Name:          "SomeStruct",
		PkgName:       "test",
		PkgImportName: "go.import.test",
		Fields:        nil,
	})

	err = tsr.WriteToDir("_test/")
	assert.NoError(t, err)

	assert.FileExists(t, "_test/github_com_widmogrod_mkunion_x_shape_testasset.ts")

	contents, err := os.ReadFile("_test/github_com_widmogrod_mkunion_x_shape_testasset.ts")
	assert.NoError(t, err)

	expected := `//generated by mkunion
export type Example = {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "testasset.A",
	"testasset.A": A
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "testasset.B",
	"testasset.B": B
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "testasset.C",
	"testasset.C": C
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "testasset.D",
	"testasset.D": D
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "testasset.E",
	"testasset.E": E
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "testasset.F",
	"testasset.F": F
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "testasset.H",
	"testasset.H": H
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "testasset.I",
	"testasset.I": I
} | {
	// $type this is optional field, that is used to enable discriminative switch-statement in TypeScript, its not part of mkunion schema
	"$type"?: "testasset.J",
	"testasset.J": J
}

export type A = {
	Name?: string,
}

export type B = {
	Age?: number,
	A?: A,
	T?: time.Time,
}

export type C = string
export type D = number
export type E = number
export type F = boolean
export type H = {[key: string]: Example}
export type I = Example[]
export type J = string[]

//eslint-disable-next-line
import * as time from './time'
`
	assert.Equal(t, expected, string(contents))
}

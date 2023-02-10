package schema

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoToSchema(t *testing.T) {
	data := AStruct{
		Foo: 123,
		Bar: 333,
	}
	expected := &Map{
		Field: []Field{
			{
				Name:  "Foo",
				Value: MkInt(123),
			}, {
				Name:  "Bar",
				Value: MkInt(333),
			},
		},
	}
	schema := FromGo(data)

	assert.Equal(
		t,
		expected,
		schema,
	)
}

func TestGoToSchemaComplex(t *testing.T) {
	someStr := "some string"

	data := BStruct{
		Foo: 123,
		Bars: []string{
			"bar",
			"baz",
		},
		Taz: map[string]string{
			"taz1": "taz2",
		},
		BaseStruct: &BaseStruct{
			Age: 123,
		},
		S: &someStr,
		List: []AStruct{
			{
				Foo: 444,
			},
		},
		Ma: map[string]AStruct{
			"key": {
				Foo: 666,
				Bar: 555,
			},
		},
	}
	expected := &Map{
		Field: []Field{
			{
				Name: "BStruct",
				Value: &Map{
					Field: []Field{
						{
							Name:  "Foo",
							Value: MkInt(123),
						}, {
							Name: "Bars",
							Value: &List{
								Items: []Schema{
									MkString("bar"),
									MkString("baz"),
								},
							},
						}, {
							Name: "Taz",
							Value: &Map{
								Field: []Field{
									{
										Name:  "taz1",
										Value: MkString("taz2"),
									},
								},
							},
						}, {
							Name: "BaseStruct",
							Value: &Map{
								Field: []Field{
									{
										Name:  "Age",
										Value: MkInt(123),
									},
								},
							},
						}, {
							Name:  "S",
							Value: MkString("some string"),
						},
						{
							Name: "List",
							Value: &List{
								Items: []Schema{
									&Map{
										Field: []Field{
											{
												Name:  "Foo",
												Value: MkInt(444),
											},
											{
												Name:  "Bar",
												Value: MkInt(0),
											},
										},
									},
								},
							},
						},
						{
							Name: "Ma",
							Value: &Map{
								Field: []Field{
									{
										Name: "key",
										Value: &Map{
											Field: []Field{
												{
													Name:  "Foo",
													Value: MkInt(666),
												},
												{
													Name:  "Bar",
													Value: MkInt(555),
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	schema := FromGo(data, WithOnlyTheseRules(
		&WrapInMap[BStruct]{InField: "BStruct"},
	))
	assert.Equal(t, expected, schema)

	result := MustToGo(schema, WithOnlyTheseRules(
		WhenPath([]string{}, UseTypeDef(&UnionMap{})),
		WhenPath([]string{"BStruct"}, UseStruct(BStruct{})),
		WhenPath([]string{"*", "BStruct", "BaseStruct"}, UseStruct(&BaseStruct{})),
		WhenPath([]string{"*", "BStruct", "List", "[*]"}, UseStruct(AStruct{})),
		WhenPath([]string{"*", "BStruct", "Ma", "key"}, UseStruct(AStruct{})),
	))
	assert.Equal(t, data, result)
}

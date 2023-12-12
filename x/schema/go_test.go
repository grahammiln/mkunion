package schema

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGoToSchema(t *testing.T) {
	data := AStruct{
		Foo: 123,
		Bar: 333,
	}
	expected := &Map{
		"Foo": MkInt(123),
		"Bar": MkInt(333),
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
		Bi: []byte("some bytes"),
		Bip: &[]byte{
			1, 2, 3, 4, 5,
		},
	}
	expected := &Map{
		"BStruct": &Map{
			"Foo": MkInt(123),
			"Bars": MkList(
				MkString("bar"),
				MkString("baz"),
			),
			"Taz": &Map{
				"taz1": MkString("taz2"),
			},
			"BaseStruct": &Map{
				"Age": MkInt(123),
			},
			"S": MkString("some string"),
			"List": MkList(
				&Map{
					"Foo": MkInt(444),
					"Bar": MkInt(0),
				},
			),
			"Ma": &Map{
				"key": &Map{
					"Foo": MkInt(666),
					"Bar": MkInt(555),
				},
			},
			"Bi":  MkBinary([]byte("some bytes")),
			"Bip": MkBinary([]byte{1, 2, 3, 4, 5}),
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

func TestToGenericGo(t *testing.T) {
	t.Run("convert with struct", func(t *testing.T) {
		data := AStruct{
			Foo: 123,
			Bar: 333,
		}
		schema := FromGo(data)
		result, err := ToGoG[AStruct](schema)
		assert.NoError(t, err)
		assert.Equal(t, data, result)
	})

	t.Run("convert with pointer", func(t *testing.T) {
		data := &AStruct{
			Foo: 123,
			Bar: 333,
		}
		schema := FromGo(data)
		result, err := ToGoG[*AStruct](schema)
		assert.NoError(t, err)
		assert.Equal(t, data, result)
	})

	t.Run("convert with interface", func(t *testing.T) {
		data := &AStruct{
			Foo: 123,
			Bar: 333,
		}
		expected := map[string]any{
			"Foo": float64(123),
			"Bar": float64(333),
		}
		schema := FromGo(data)
		result, err := ToGoG[any](schema)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("convert primitive type", func(t *testing.T) {
		data := 123
		schema := FromGo(data)
		result, err := ToGoG[int](schema)
		assert.NoError(t, err)
		assert.Equal(t, data, result)
	})
}

type ExternalPackageStruct struct {
	Duration     time.Duration            `desc:"time duration"`
	DurationList []time.Duration          `desc:"time duration list"`
	DurationMap  map[string]time.Duration `desc:"time duration map"`

	DurationRef *time.Duration `desc:"time duration ref"`

	Time     time.Time   `desc:"time"`
	TimeList []time.Time `desc:"time list"`
	TimeMap  map[string]time.Time
	TimeRef  *time.Time `desc:"time ref"`
}

func TestExternalPackageSerDe(t *testing.T) {
	now, _ := time.Parse(time.RFC3339Nano, time.Now().Format(time.RFC3339Nano))
	hour := time.Hour
	data := ExternalPackageStruct{
		Duration: time.Second,
		DurationList: []time.Duration{
			time.Second,
			time.Minute,
		},
		DurationMap: map[string]time.Duration{
			"hour": time.Hour,
		},
		DurationRef: &hour,
		Time:        now,
		TimeList: []time.Time{
			now,
			now.Add(time.Hour),
		},
		TimeMap: map[string]time.Time{
			"now":    now,
			"now+1h": now.Add(time.Hour),
		},
	}

	schema := FromGo(data)
	result, err := ToGoG[ExternalPackageStruct](schema)
	assert.NoError(t, err)
	assert.Equal(t, data, result)
}

package mkunion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunctionMatchGenerator_Generate(t *testing.T) {
	g := FunctionMatchGenerator{
		Header:      Header,
		PackageName: "f",
		MaxSize:     3,
	}

	result, err := g.Generate()
	assert.NoError(t, err)
	assert.Equal(t, `// Code generated by mkunion. DO NOT EDIT.
package f

import (
	"errors"
	"fmt"
)

func Match2[TIn, TOut, T1, T2 any](
	x TIn,
	f1 func(x T1) TOut,
	f2 func(x T2) TOut,
	df func(x TIn) TOut,
) TOut {
	switch y := any(x).(type) {
	case T1:
		return f1(y)
	case T2:
		return f2(y)
	}

	return df(x)
}

func MustMatch2[TIn, TOut, T1, T2 any](
	x TIn,
	f1 func(x T1) TOut,
	f2 func(x T2) TOut,
) TOut {
	return Match2(x, f1, f2, func(x TIn) TOut {
		var t1 T1
		var t2 T2
		panic(errors.New(fmt.Sprintf("unexpected match type %T. expected (%T or %T)", x, t1, t2)))
	})
}

func Match3[TIn, TOut, T1, T2, T3 any](
	x TIn,
	f1 func(x T1) TOut,
	f2 func(x T2) TOut,
	f3 func(x T3) TOut,
	df func(x TIn) TOut,
) TOut {
	switch y := any(x).(type) {
	case T1:
		return f1(y)
	case T2:
		return f2(y)
	case T3:
		return f3(y)
	}

	return df(x)
}

func MustMatch3[TIn, TOut, T1, T2, T3 any](
	x TIn,
	f1 func(x T1) TOut,
	f2 func(x T2) TOut,
	f3 func(x T3) TOut,
) TOut {
	return Match3(x, f1, f2, f3, func(x TIn) TOut {
		var t1 T1
		var t2 T2
		var t3 T3
		panic(errors.New(fmt.Sprintf("unexpected match type %T. expected (%T or %T or %T)", x, t1, t2, t3)))
	})
}
`, string(result))
}

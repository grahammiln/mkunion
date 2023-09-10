package schema

import (
	"bytes"
	"strconv"
	"strings"
)

func As[A int | int8 | int16 | int32 | int64 |
	uint | uint8 | uint16 | uint32 | uint64 |
	float32 | float64 |
	bool | string | []byte](x Schema) (A, bool) {
	var def A
	if x == nil {
		if any(def) == nil {
			return def, true
		}
		return def, false
	}

	return MustMatchSchemaR2(
		x,
		func(x *None) (A, bool) {
			if any(def) == nil {
				return def, true
			}
			return def, false
		},
		func(x *Bool) (A, bool) {
			switch any(def).(type) {
			case bool:
				return any(bool(*x)).(A), true
			}

			return def, false
		},
		func(x *Number) (A, bool) {
			switch any(def).(type) {
			case float32:
				return any(float32(*x)).(A), true
			case float64:
				return any(float64(*x)).(A), true
			case int:
				return any(int(*x)).(A), true
			case int8:
				return any(int8(*x)).(A), true
			case int16:
				return any(int16(*x)).(A), true
			case int32:
				return any(int32(*x)).(A), true
			case int64:
				return any(int64(*x)).(A), true
			case uint:
				return any(uint(*x)).(A), true
			case uint8:
				return any(uint8(*x)).(A), true
			case uint16:
				return any(uint16(*x)).(A), true
			case uint32:
				return any(uint32(*x)).(A), true
			case uint64:
				return any(uint64(*x)).(A), true
			}
			return def, false
		},
		func(x *String) (A, bool) {
			switch any(def).(type) {
			case string:
				return any(string(*x)).(A), true
			case []byte:
				return any([]byte(*x)).(A), true
			case float64:
				v, err := strconv.ParseFloat(string(*x), 64)
				if err != nil {
					return def, false
				}
				return any(v).(A), true
			case float32:
				v, err := strconv.ParseFloat(string(*x), 32)
				if err != nil {
					return def, false
				}
				return any(float32(v)).(A), true
			case int:
				v, err := strconv.Atoi(string(*x))
				if err != nil {
					return def, false
				}
				return any(v).(A), true
			case int8:
				v, err := strconv.ParseInt(string(*x), 10, 8)
				if err != nil {
					return def, false
				}
				return any(int8(v)).(A), true
			case int16:
				v, err := strconv.ParseInt(string(*x), 10, 16)
				if err != nil {
					return def, false
				}
				return any(int16(v)).(A), true
			case int32:
				v, err := strconv.ParseInt(string(*x), 10, 32)
				if err != nil {
					return def, false
				}
				return any(int32(v)).(A), true
			case int64:
				v, err := strconv.ParseInt(string(*x), 10, 64)
				if err != nil {
					return def, false
				}
				return any(int64(v)).(A), true
			}

			return def, false
		},
		func(x *Binary) (A, bool) {
			switch any(def).(type) {
			case []byte:
				return any(x.B).(A), true
			case string:
				return any(string(x.B)).(A), true
			}

			return def, false
		},
		func(x *List) (A, bool) {
			return def, false
		},
		func(x *Map) (A, bool) {
			return def, false
		})
}

func AsDefault[A int | int8 | int16 | int32 | int64 |
	uint | uint8 | uint16 | uint32 | uint64 |
	float32 | float64 |
	bool | string | []byte](x Schema, def A) A {

	res, ok := As[A](x)
	if ok {
		return res
	}

	return def
}

func LocationToPath(location string) []string {
	var result []string

	var value []string
	var isQuote bool = false
	var isLit bool = false
	for _, char := range strings.Split(location, "") {
		switch char {
		case "'":
			isQuote = !isQuote

		case ".":
			if isLit && !isQuote {
				isLit = false
				result = append(result, strings.Join(value, ""))
				value = nil

			} else {
				if isQuote {
					value = append(value, char)
				} else {
					result = append(result, strings.Join(value, ""))
					value = nil
				}
			}

		default:
			if !isLit && !isQuote {
				isLit = true
			}

			if isQuote || isLit {
				value = append(value, char)
			} else {
				result = append(result, strings.Join(value, ""))
				value = nil
			}
		}
	}

	if len(value) > 0 {
		result = append(result, strings.Join(value, ""))
	}

	return result
}

func Get(data Schema, location string) Schema {
	path := LocationToPath(location)
	for _, p := range path {
		if p == "" {
			return nil
		}

		if strings.HasPrefix(p, "[") && strings.HasSuffix(p, "]") {
			idx := strings.TrimPrefix(p, "[")
			idx = strings.TrimSuffix(idx, "]")
			i, err := strconv.Atoi(idx)
			if err != nil {
				return nil
			}

			listData, ok := data.(*List)
			if ok && len(listData.Items) > i {
				data = listData.Items[i]
				continue
			}

			return nil
		}

		mapData, ok := data.(*Map)
		if !ok {
			return nil
		}

		if p == "#" && len(mapData.Field) == 1 {
			data = mapData.Field[0].Value
			continue
		}

		var found bool
		for _, item := range mapData.Field {
			if item.Name == p {
				found = true
				data = item.Value
				break
			}
		}

		if !found {
			return nil
		}
	}

	return data
}

func Reduce[A any](data Schema, init A, fn func(Schema, A) A) A {
	if data == nil {
		return init
	}

	return MustMatchSchema(
		data,
		func(x *None) A {
			return init
		},
		func(x *Bool) A {
			return fn(x, init)
		},
		func(x *Number) A {
			return fn(x, init)
		},
		func(x *String) A {
			return fn(x, init)
		},
		func(x *Binary) A {
			return fn(x, init)
		},
		func(x *List) A {
			for _, y := range x.Items {
				init = fn(y, init)
			}

			return init
		},
		func(x *Map) A {
			for _, y := range x.Field {
				init = fn(y.Value, init)
			}

			return init
		},
	)
}

func Compare(a, b Schema) int {
	if a == nil {
		a = none
	}
	if b == nil {
		b = none
	}

	return MustMatchSchema(
		a,
		func(x *None) int {
			switch b.(type) {
			case *None:
				return 0
			}

			return -1
		},
		func(x *Bool) int {
			switch y := b.(type) {
			case *None:
				return 1
			case *Bool:
				if *x == *y {
					return 0
				}
				if *x {
					return 1
				}
				return -1
			}

			return -1
		},
		func(x *Number) int {
			switch y := b.(type) {
			case *None, *Bool:
				return 1
			case *Number:
				return int(*x - *y)
			}

			return -1
		},
		func(x *String) int {
			switch y := b.(type) {
			case *None, *Bool, *Number:
				return 1
			case *String:
				return strings.Compare(string(*x), string(*y))
			}

			return -1
		},
		func(x *Binary) int {
			switch y := b.(type) {
			case *None, *Bool, *Number, *String:
				return 1
			case *Binary:
				return bytes.Compare(x.B, y.B)
			}

			return -1
		},
		func(x *List) int {
			switch y := b.(type) {
			case *None, *Bool, *Number, *String, *Binary:
				return 1
			case *List:
				if len(x.Items) == len(y.Items) {
					for i := range x.Items {
						cmp := Compare(x.Items[i], y.Items[i])
						if cmp != 0 {
							return cmp
						}
					}
					return 0
				}
				if len(x.Items) > len(y.Items) {
					return 1
				}

				return -1
			}

			return -1

		},
		func(x *Map) int {
			switch y := b.(type) {
			case *None, *Bool, *Number, *String, *Binary, *List:
				return 1
			case *Map:
				if len(x.Field) == len(y.Field) {
					for _, xField := range x.Field {
						var found bool
						for _, yField := range y.Field {
							if yField.Name == xField.Name {
								found = true
								cmp := Compare(xField.Value, yField.Value)
								if cmp != 0 {
									return cmp
								}
								break
							}
						}
						if !found {
							return -1
						}
					}
					return 0
				}

				if len(x.Field) > len(y.Field) {
					return 1
				}

				return -1
			}

			return -1
		},
	)
}

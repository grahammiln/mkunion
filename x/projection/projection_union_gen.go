// Code generated by mkunion. DO NOT EDIT.
package projection

import (
	"encoding/json"
	"fmt"
	"github.com/widmogrod/mkunion/x/shared"
)

type DataVisitor[A any] interface {
	VisitRecord(v *Record[A]) any
	VisitWatermark(v *Watermark[A]) any
}

type Data[A any] interface {
	AcceptData(g DataVisitor[A]) any
}

var (
	_ Data[any] = (*Record[any])(nil)
	_ Data[any] = (*Watermark[any])(nil)
)

func (r *Record[A]) AcceptData(v DataVisitor[A]) any    { return v.VisitRecord(r) }
func (r *Watermark[A]) AcceptData(v DataVisitor[A]) any { return v.VisitWatermark(r) }

func MatchDataR3[A any, T0, T1, T2 any](
	x Data[A],
	f1 func(x *Record[A]) (T0, T1, T2),
	f2 func(x *Watermark[A]) (T0, T1, T2),
) (T0, T1, T2) {
	switch v := x.(type) {
	case *Record[A]:
		return f1(v)
	case *Watermark[A]:
		return f2(v)
	}
	var result1 T0
	var result2 T1
	var result3 T2
	return result1, result2, result3
}

func MatchDataR2[A any, T0, T1 any](
	x Data[A],
	f1 func(x *Record[A]) (T0, T1),
	f2 func(x *Watermark[A]) (T0, T1),
) (T0, T1) {
	switch v := x.(type) {
	case *Record[A]:
		return f1(v)
	case *Watermark[A]:
		return f2(v)
	}
	var result1 T0
	var result2 T1
	return result1, result2
}

func MatchDataR1[A any, T0 any](
	x Data[A],
	f1 func(x *Record[A]) T0,
	f2 func(x *Watermark[A]) T0,
) T0 {
	switch v := x.(type) {
	case *Record[A]:
		return f1(v)
	case *Watermark[A]:
		return f2(v)
	}
	var result1 T0
	return result1
}

func MatchDataR0[A any](
	x Data[A],
	f1 func(x *Record[A]),
	f2 func(x *Watermark[A]),
) {
	switch v := x.(type) {
	case *Record[A]:
		f1(v)
	case *Watermark[A]:
		f2(v)
	}
}
func init() {
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/projection.Data[any]", DataFromJSON[any], DataToJSON[any])
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/projection.Record[any]", RecordFromJSON[any], RecordToJSON[any])
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/projection.Watermark[any]", WatermarkFromJSON[any], WatermarkToJSON[any])
}

type DataUnionJSON[A any] struct {
	Type      string          `json:"$type,omitempty"`
	Record    json.RawMessage `json:"projection.Record,omitempty"`
	Watermark json.RawMessage `json:"projection.Watermark,omitempty"`
}

func DataFromJSON[A any](x []byte) (Data[A], error) {
	if x == nil || len(x) == 0 {
		return nil, nil
	}
	if string(x[:4]) == "null" {
		return nil, nil
	}
	var data DataUnionJSON[A]
	err := json.Unmarshal(x, &data)
	if err != nil {
		return nil, fmt.Errorf("projection.DataFromJSON[A]: %w", err)
	}

	switch data.Type {
	case "projection.Record":
		return RecordFromJSON[A](data.Record)
	case "projection.Watermark":
		return WatermarkFromJSON[A](data.Watermark)
	}

	if data.Record != nil {
		return RecordFromJSON[A](data.Record)
	} else if data.Watermark != nil {
		return WatermarkFromJSON[A](data.Watermark)
	}
	return nil, fmt.Errorf("projection.DataFromJSON[A]: unknown type: %s", data.Type)
}

func DataToJSON[A any](x Data[A]) ([]byte, error) {
	if x == nil {
		return []byte(`null`), nil
	}
	return MatchDataR2(
		x,
		func(y *Record[A]) ([]byte, error) {
			body, err := RecordToJSON[A](y)
			if err != nil {
				return nil, fmt.Errorf("projection.DataToJSON[A]: %w", err)
			}
			return json.Marshal(DataUnionJSON[A]{
				Type:   "projection.Record",
				Record: body,
			})
		},
		func(y *Watermark[A]) ([]byte, error) {
			body, err := WatermarkToJSON[A](y)
			if err != nil {
				return nil, fmt.Errorf("projection.DataToJSON[A]: %w", err)
			}
			return json.Marshal(DataUnionJSON[A]{
				Type:      "projection.Watermark",
				Watermark: body,
			})
		},
	)
}

func RecordFromJSON[A any](x []byte) (*Record[A], error) {
	result := new(Record[A])
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, fmt.Errorf("projection.RecordFromJSON[A]: %w", err)
	}
	return result, nil
}

func RecordToJSON[A any](x *Record[A]) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*Record[any])(nil)
	_ json.Marshaler   = (*Record[any])(nil)
)

func (r *Record[A]) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONRecordLb_A_bL(*r)
}
func (r *Record[A]) _marshalJSONRecordLb_A_bL(x Record[A]) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldKey []byte
	fieldKey, err = r._marshalJSONstring(x.Key)
	if err != nil {
		return nil, fmt.Errorf("projection: Record[A]._marshalJSONRecordLb_A_bL: field name Key; %w", err)
	}
	partial["Key"] = fieldKey
	var fieldData []byte
	fieldData, err = r._marshalJSONA(x.Data)
	if err != nil {
		return nil, fmt.Errorf("projection: Record[A]._marshalJSONRecordLb_A_bL: field name Data; %w", err)
	}
	partial["Data"] = fieldData
	var fieldEventTime []byte
	fieldEventTime, err = r._marshalJSONEventTime(x.EventTime)
	if err != nil {
		return nil, fmt.Errorf("projection: Record[A]._marshalJSONRecordLb_A_bL: field name EventTime; %w", err)
	}
	partial["EventTime"] = fieldEventTime
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("projection: Record[A]._marshalJSONRecordLb_A_bL: struct; %w", err)
	}
	return result, nil
}
func (r *Record[A]) _marshalJSONstring(x string) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("projection: Record[A]._marshalJSONstring:; %w", err)
	}
	return result, nil
}
func (r *Record[A]) _marshalJSONA(x A) ([]byte, error) {
	result, err := shared.JSONMarshal[A](x)
	if err != nil {
		return nil, fmt.Errorf("projection: Record[A]._marshalJSONA:; %w", err)
	}
	return result, nil
}
func (r *Record[A]) _marshalJSONEventTime(x EventTime) ([]byte, error) {
	result, err := shared.JSONMarshal[EventTime](x)
	if err != nil {
		return nil, fmt.Errorf("projection: Record[A]._marshalJSONEventTime:; %w", err)
	}
	return result, nil
}
func (r *Record[A]) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONRecordLb_A_bL(data)
	if err != nil {
		return fmt.Errorf("projection: Record[A].UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *Record[A]) _unmarshalJSONRecordLb_A_bL(data []byte) (Record[A], error) {
	result := Record[A]{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("projection: Record[A]._unmarshalJSONRecordLb_A_bL: native struct unwrap; %w", err)
	}
	if fieldKey, ok := partial["Key"]; ok {
		result.Key, err = r._unmarshalJSONstring(fieldKey)
		if err != nil {
			return result, fmt.Errorf("projection: Record[A]._unmarshalJSONRecordLb_A_bL: field Key; %w", err)
		}
	}
	if fieldData, ok := partial["Data"]; ok {
		result.Data, err = r._unmarshalJSONA(fieldData)
		if err != nil {
			return result, fmt.Errorf("projection: Record[A]._unmarshalJSONRecordLb_A_bL: field Data; %w", err)
		}
	}
	if fieldEventTime, ok := partial["EventTime"]; ok {
		result.EventTime, err = r._unmarshalJSONEventTime(fieldEventTime)
		if err != nil {
			return result, fmt.Errorf("projection: Record[A]._unmarshalJSONRecordLb_A_bL: field EventTime; %w", err)
		}
	}
	return result, nil
}
func (r *Record[A]) _unmarshalJSONstring(data []byte) (string, error) {
	var result string
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("projection: Record[A]._unmarshalJSONstring: native primitive unwrap; %w", err)
	}
	return result, nil
}
func (r *Record[A]) _unmarshalJSONA(data []byte) (A, error) {
	result, err := shared.JSONUnmarshal[A](data)
	if err != nil {
		return result, fmt.Errorf("projection: Record[A]._unmarshalJSONA: native ref unwrap; %w", err)
	}
	return result, nil
}
func (r *Record[A]) _unmarshalJSONEventTime(data []byte) (EventTime, error) {
	result, err := shared.JSONUnmarshal[EventTime](data)
	if err != nil {
		return result, fmt.Errorf("projection: Record[A]._unmarshalJSONEventTime: native ref unwrap; %w", err)
	}
	return result, nil
}

func WatermarkFromJSON[A any](x []byte) (*Watermark[A], error) {
	result := new(Watermark[A])
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, fmt.Errorf("projection.WatermarkFromJSON[A]: %w", err)
	}
	return result, nil
}

func WatermarkToJSON[A any](x *Watermark[A]) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*Watermark[any])(nil)
	_ json.Marshaler   = (*Watermark[any])(nil)
)

func (r *Watermark[A]) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONWatermarkLb_A_bL(*r)
}
func (r *Watermark[A]) _marshalJSONWatermarkLb_A_bL(x Watermark[A]) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldKey []byte
	fieldKey, err = r._marshalJSONstring(x.Key)
	if err != nil {
		return nil, fmt.Errorf("projection: Watermark[A]._marshalJSONWatermarkLb_A_bL: field name Key; %w", err)
	}
	partial["Key"] = fieldKey
	var fieldEventTime []byte
	fieldEventTime, err = r._marshalJSONEventTime(x.EventTime)
	if err != nil {
		return nil, fmt.Errorf("projection: Watermark[A]._marshalJSONWatermarkLb_A_bL: field name EventTime; %w", err)
	}
	partial["EventTime"] = fieldEventTime
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("projection: Watermark[A]._marshalJSONWatermarkLb_A_bL: struct; %w", err)
	}
	return result, nil
}
func (r *Watermark[A]) _marshalJSONstring(x string) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("projection: Watermark[A]._marshalJSONstring:; %w", err)
	}
	return result, nil
}
func (r *Watermark[A]) _marshalJSONEventTime(x EventTime) ([]byte, error) {
	result, err := shared.JSONMarshal[EventTime](x)
	if err != nil {
		return nil, fmt.Errorf("projection: Watermark[A]._marshalJSONEventTime:; %w", err)
	}
	return result, nil
}
func (r *Watermark[A]) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONWatermarkLb_A_bL(data)
	if err != nil {
		return fmt.Errorf("projection: Watermark[A].UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *Watermark[A]) _unmarshalJSONWatermarkLb_A_bL(data []byte) (Watermark[A], error) {
	result := Watermark[A]{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("projection: Watermark[A]._unmarshalJSONWatermarkLb_A_bL: native struct unwrap; %w", err)
	}
	if fieldKey, ok := partial["Key"]; ok {
		result.Key, err = r._unmarshalJSONstring(fieldKey)
		if err != nil {
			return result, fmt.Errorf("projection: Watermark[A]._unmarshalJSONWatermarkLb_A_bL: field Key; %w", err)
		}
	}
	if fieldEventTime, ok := partial["EventTime"]; ok {
		result.EventTime, err = r._unmarshalJSONEventTime(fieldEventTime)
		if err != nil {
			return result, fmt.Errorf("projection: Watermark[A]._unmarshalJSONWatermarkLb_A_bL: field EventTime; %w", err)
		}
	}
	return result, nil
}
func (r *Watermark[A]) _unmarshalJSONstring(data []byte) (string, error) {
	var result string
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("projection: Watermark[A]._unmarshalJSONstring: native primitive unwrap; %w", err)
	}
	return result, nil
}
func (r *Watermark[A]) _unmarshalJSONEventTime(data []byte) (EventTime, error) {
	result, err := shared.JSONUnmarshal[EventTime](data)
	if err != nil {
		return result, fmt.Errorf("projection: Watermark[A]._unmarshalJSONEventTime: native ref unwrap; %w", err)
	}
	return result, nil
}

type EitherVisitor[A any, B any] interface {
	VisitLeft(v *Left[A, B]) any
	VisitRight(v *Right[A, B]) any
}

type Either[A any, B any] interface {
	AcceptEither(g EitherVisitor[A, B]) any
}

var (
	_ Either[any, any] = (*Left[any, any])(nil)
	_ Either[any, any] = (*Right[any, any])(nil)
)

func (r *Left[A, B]) AcceptEither(v EitherVisitor[A, B]) any  { return v.VisitLeft(r) }
func (r *Right[A, B]) AcceptEither(v EitherVisitor[A, B]) any { return v.VisitRight(r) }

func MatchEitherR3[A any, B any, T0, T1, T2 any](
	x Either[A, B],
	f1 func(x *Left[A, B]) (T0, T1, T2),
	f2 func(x *Right[A, B]) (T0, T1, T2),
) (T0, T1, T2) {
	switch v := x.(type) {
	case *Left[A, B]:
		return f1(v)
	case *Right[A, B]:
		return f2(v)
	}
	var result1 T0
	var result2 T1
	var result3 T2
	return result1, result2, result3
}

func MatchEitherR2[A any, B any, T0, T1 any](
	x Either[A, B],
	f1 func(x *Left[A, B]) (T0, T1),
	f2 func(x *Right[A, B]) (T0, T1),
) (T0, T1) {
	switch v := x.(type) {
	case *Left[A, B]:
		return f1(v)
	case *Right[A, B]:
		return f2(v)
	}
	var result1 T0
	var result2 T1
	return result1, result2
}

func MatchEitherR1[A any, B any, T0 any](
	x Either[A, B],
	f1 func(x *Left[A, B]) T0,
	f2 func(x *Right[A, B]) T0,
) T0 {
	switch v := x.(type) {
	case *Left[A, B]:
		return f1(v)
	case *Right[A, B]:
		return f2(v)
	}
	var result1 T0
	return result1
}

func MatchEitherR0[A any, B any](
	x Either[A, B],
	f1 func(x *Left[A, B]),
	f2 func(x *Right[A, B]),
) {
	switch v := x.(type) {
	case *Left[A, B]:
		f1(v)
	case *Right[A, B]:
		f2(v)
	}
}
func init() {
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/projection.Either[any,any]", EitherFromJSON[any, any], EitherToJSON[any, any])
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/projection.Left[any,any]", LeftFromJSON[any, any], LeftToJSON[any, any])
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/projection.Right[any,any]", RightFromJSON[any, any], RightToJSON[any, any])
}

type EitherUnionJSON[A any, B any] struct {
	Type  string          `json:"$type,omitempty"`
	Left  json.RawMessage `json:"projection.Left,omitempty"`
	Right json.RawMessage `json:"projection.Right,omitempty"`
}

func EitherFromJSON[A any, B any](x []byte) (Either[A, B], error) {
	if x == nil || len(x) == 0 {
		return nil, nil
	}
	if string(x[:4]) == "null" {
		return nil, nil
	}
	var data EitherUnionJSON[A, B]
	err := json.Unmarshal(x, &data)
	if err != nil {
		return nil, fmt.Errorf("projection.EitherFromJSON[A,B]: %w", err)
	}

	switch data.Type {
	case "projection.Left":
		return LeftFromJSON[A, B](data.Left)
	case "projection.Right":
		return RightFromJSON[A, B](data.Right)
	}

	if data.Left != nil {
		return LeftFromJSON[A, B](data.Left)
	} else if data.Right != nil {
		return RightFromJSON[A, B](data.Right)
	}
	return nil, fmt.Errorf("projection.EitherFromJSON[A,B]: unknown type: %s", data.Type)
}

func EitherToJSON[A any, B any](x Either[A, B]) ([]byte, error) {
	if x == nil {
		return []byte(`null`), nil
	}
	return MatchEitherR2(
		x,
		func(y *Left[A, B]) ([]byte, error) {
			body, err := LeftToJSON[A, B](y)
			if err != nil {
				return nil, fmt.Errorf("projection.EitherToJSON[A,B]: %w", err)
			}
			return json.Marshal(EitherUnionJSON[A, B]{
				Type: "projection.Left",
				Left: body,
			})
		},
		func(y *Right[A, B]) ([]byte, error) {
			body, err := RightToJSON[A, B](y)
			if err != nil {
				return nil, fmt.Errorf("projection.EitherToJSON[A,B]: %w", err)
			}
			return json.Marshal(EitherUnionJSON[A, B]{
				Type:  "projection.Right",
				Right: body,
			})
		},
	)
}

func LeftFromJSON[A any, B any](x []byte) (*Left[A, B], error) {
	result := new(Left[A, B])
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, fmt.Errorf("projection.LeftFromJSON[A,B]: %w", err)
	}
	return result, nil
}

func LeftToJSON[A any, B any](x *Left[A, B]) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*Left[any, any])(nil)
	_ json.Marshaler   = (*Left[any, any])(nil)
)

func (r *Left[A, B]) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONLeftLb_ACommaB_bL(*r)
}
func (r *Left[A, B]) _marshalJSONLeftLb_ACommaB_bL(x Left[A, B]) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldLeft []byte
	fieldLeft, err = r._marshalJSONA(x.Left)
	if err != nil {
		return nil, fmt.Errorf("projection: Left[A,B]._marshalJSONLeftLb_ACommaB_bL: field name Left; %w", err)
	}
	partial["Left"] = fieldLeft
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("projection: Left[A,B]._marshalJSONLeftLb_ACommaB_bL: struct; %w", err)
	}
	return result, nil
}
func (r *Left[A, B]) _marshalJSONA(x A) ([]byte, error) {
	result, err := shared.JSONMarshal[A](x)
	if err != nil {
		return nil, fmt.Errorf("projection: Left[A,B]._marshalJSONA:; %w", err)
	}
	return result, nil
}
func (r *Left[A, B]) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONLeftLb_ACommaB_bL(data)
	if err != nil {
		return fmt.Errorf("projection: Left[A,B].UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *Left[A, B]) _unmarshalJSONLeftLb_ACommaB_bL(data []byte) (Left[A, B], error) {
	result := Left[A, B]{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("projection: Left[A,B]._unmarshalJSONLeftLb_ACommaB_bL: native struct unwrap; %w", err)
	}
	if fieldLeft, ok := partial["Left"]; ok {
		result.Left, err = r._unmarshalJSONA(fieldLeft)
		if err != nil {
			return result, fmt.Errorf("projection: Left[A,B]._unmarshalJSONLeftLb_ACommaB_bL: field Left; %w", err)
		}
	}
	return result, nil
}
func (r *Left[A, B]) _unmarshalJSONA(data []byte) (A, error) {
	result, err := shared.JSONUnmarshal[A](data)
	if err != nil {
		return result, fmt.Errorf("projection: Left[A,B]._unmarshalJSONA: native ref unwrap; %w", err)
	}
	return result, nil
}

func RightFromJSON[A any, B any](x []byte) (*Right[A, B], error) {
	result := new(Right[A, B])
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, fmt.Errorf("projection.RightFromJSON[A,B]: %w", err)
	}
	return result, nil
}

func RightToJSON[A any, B any](x *Right[A, B]) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*Right[any, any])(nil)
	_ json.Marshaler   = (*Right[any, any])(nil)
)

func (r *Right[A, B]) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONRightLb_ACommaB_bL(*r)
}
func (r *Right[A, B]) _marshalJSONRightLb_ACommaB_bL(x Right[A, B]) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldRight []byte
	fieldRight, err = r._marshalJSONB(x.Right)
	if err != nil {
		return nil, fmt.Errorf("projection: Right[A,B]._marshalJSONRightLb_ACommaB_bL: field name Right; %w", err)
	}
	partial["Right"] = fieldRight
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("projection: Right[A,B]._marshalJSONRightLb_ACommaB_bL: struct; %w", err)
	}
	return result, nil
}
func (r *Right[A, B]) _marshalJSONB(x B) ([]byte, error) {
	result, err := shared.JSONMarshal[B](x)
	if err != nil {
		return nil, fmt.Errorf("projection: Right[A,B]._marshalJSONB:; %w", err)
	}
	return result, nil
}
func (r *Right[A, B]) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONRightLb_ACommaB_bL(data)
	if err != nil {
		return fmt.Errorf("projection: Right[A,B].UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *Right[A, B]) _unmarshalJSONRightLb_ACommaB_bL(data []byte) (Right[A, B], error) {
	result := Right[A, B]{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("projection: Right[A,B]._unmarshalJSONRightLb_ACommaB_bL: native struct unwrap; %w", err)
	}
	if fieldRight, ok := partial["Right"]; ok {
		result.Right, err = r._unmarshalJSONB(fieldRight)
		if err != nil {
			return result, fmt.Errorf("projection: Right[A,B]._unmarshalJSONRightLb_ACommaB_bL: field Right; %w", err)
		}
	}
	return result, nil
}
func (r *Right[A, B]) _unmarshalJSONB(data []byte) (B, error) {
	result, err := shared.JSONUnmarshal[B](data)
	if err != nil {
		return result, fmt.Errorf("projection: Right[A,B]._unmarshalJSONB: native ref unwrap; %w", err)
	}
	return result, nil
}

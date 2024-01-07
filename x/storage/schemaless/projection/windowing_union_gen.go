// Code generated by mkunion. DO NOT EDIT.
package projection

import (
	"encoding/json"
	"fmt"
	"github.com/widmogrod/mkunion/x/shared"
	"time"
)

type WindowDescriptionVisitor interface {
	VisitSessionWindow(v *SessionWindow) any
	VisitSlidingWindow(v *SlidingWindow) any
	VisitFixedWindow(v *FixedWindow) any
}

type WindowDescription interface {
	AcceptWindowDescription(g WindowDescriptionVisitor) any
}

var (
	_ WindowDescription = (*SessionWindow)(nil)
	_ WindowDescription = (*SlidingWindow)(nil)
	_ WindowDescription = (*FixedWindow)(nil)
)

func (r *SessionWindow) AcceptWindowDescription(v WindowDescriptionVisitor) any {
	return v.VisitSessionWindow(r)
}
func (r *SlidingWindow) AcceptWindowDescription(v WindowDescriptionVisitor) any {
	return v.VisitSlidingWindow(r)
}
func (r *FixedWindow) AcceptWindowDescription(v WindowDescriptionVisitor) any {
	return v.VisitFixedWindow(r)
}

func MatchWindowDescriptionR3[T0, T1, T2 any](
	x WindowDescription,
	f1 func(x *SessionWindow) (T0, T1, T2),
	f2 func(x *SlidingWindow) (T0, T1, T2),
	f3 func(x *FixedWindow) (T0, T1, T2),
) (T0, T1, T2) {
	switch v := x.(type) {
	case *SessionWindow:
		return f1(v)
	case *SlidingWindow:
		return f2(v)
	case *FixedWindow:
		return f3(v)
	}
	var result1 T0
	var result2 T1
	var result3 T2
	return result1, result2, result3
}

func MatchWindowDescriptionR2[T0, T1 any](
	x WindowDescription,
	f1 func(x *SessionWindow) (T0, T1),
	f2 func(x *SlidingWindow) (T0, T1),
	f3 func(x *FixedWindow) (T0, T1),
) (T0, T1) {
	switch v := x.(type) {
	case *SessionWindow:
		return f1(v)
	case *SlidingWindow:
		return f2(v)
	case *FixedWindow:
		return f3(v)
	}
	var result1 T0
	var result2 T1
	return result1, result2
}

func MatchWindowDescriptionR1[T0 any](
	x WindowDescription,
	f1 func(x *SessionWindow) T0,
	f2 func(x *SlidingWindow) T0,
	f3 func(x *FixedWindow) T0,
) T0 {
	switch v := x.(type) {
	case *SessionWindow:
		return f1(v)
	case *SlidingWindow:
		return f2(v)
	case *FixedWindow:
		return f3(v)
	}
	var result1 T0
	return result1
}

func MatchWindowDescriptionR0(
	x WindowDescription,
	f1 func(x *SessionWindow),
	f2 func(x *SlidingWindow),
	f3 func(x *FixedWindow),
) {
	switch v := x.(type) {
	case *SessionWindow:
		f1(v)
	case *SlidingWindow:
		f2(v)
	case *FixedWindow:
		f3(v)
	}
}
func init() {
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/storage/schemaless/projection.WindowDescription", WindowDescriptionFromJSON, WindowDescriptionToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/storage/schemaless/projection.SessionWindow", SessionWindowFromJSON, SessionWindowToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/storage/schemaless/projection.SlidingWindow", SlidingWindowFromJSON, SlidingWindowToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/storage/schemaless/projection.FixedWindow", FixedWindowFromJSON, FixedWindowToJSON)
}

type WindowDescriptionUnionJSON struct {
	Type          string          `json:"$type,omitempty"`
	SessionWindow json.RawMessage `json:"projection.SessionWindow,omitempty"`
	SlidingWindow json.RawMessage `json:"projection.SlidingWindow,omitempty"`
	FixedWindow   json.RawMessage `json:"projection.FixedWindow,omitempty"`
}

func WindowDescriptionFromJSON(x []byte) (WindowDescription, error) {
	if x == nil || len(x) == 0 {
		return nil, nil
	}
	if string(x[:4]) == "null" {
		return nil, nil
	}

	var data WindowDescriptionUnionJSON
	err := json.Unmarshal(x, &data)
	if err != nil {
		return nil, err
	}

	switch data.Type {
	case "projection.SessionWindow":
		return SessionWindowFromJSON(data.SessionWindow)
	case "projection.SlidingWindow":
		return SlidingWindowFromJSON(data.SlidingWindow)
	case "projection.FixedWindow":
		return FixedWindowFromJSON(data.FixedWindow)
	}

	if data.SessionWindow != nil {
		return SessionWindowFromJSON(data.SessionWindow)
	} else if data.SlidingWindow != nil {
		return SlidingWindowFromJSON(data.SlidingWindow)
	} else if data.FixedWindow != nil {
		return FixedWindowFromJSON(data.FixedWindow)
	}

	return nil, fmt.Errorf("projection.WindowDescription: unknown type %s", data.Type)
}

func WindowDescriptionToJSON(x WindowDescription) ([]byte, error) {
	if x == nil {
		return nil, nil
	}
	return MatchWindowDescriptionR2(
		x,
		func(x *SessionWindow) ([]byte, error) {
			body, err := SessionWindowToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(WindowDescriptionUnionJSON{
				Type:          "projection.SessionWindow",
				SessionWindow: body,
			})
		},
		func(x *SlidingWindow) ([]byte, error) {
			body, err := SlidingWindowToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(WindowDescriptionUnionJSON{
				Type:          "projection.SlidingWindow",
				SlidingWindow: body,
			})
		},
		func(x *FixedWindow) ([]byte, error) {
			body, err := FixedWindowToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(WindowDescriptionUnionJSON{
				Type:        "projection.FixedWindow",
				FixedWindow: body,
			})
		},
	)
}

func SessionWindowFromJSON(x []byte) (*SessionWindow, error) {
	result := new(SessionWindow)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func SessionWindowToJSON(x *SessionWindow) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*SessionWindow)(nil)
	_ json.Marshaler   = (*SessionWindow)(nil)
)

func (r *SessionWindow) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONSessionWindow(*r)
}
func (r *SessionWindow) _marshalJSONSessionWindow(x SessionWindow) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldGapDuration []byte
	fieldGapDuration, err = r._marshalJSONtime_Duration(x.GapDuration)
	if err != nil {
		return nil, fmt.Errorf("projection: SessionWindow._marshalJSONSessionWindow: field name GapDuration; %w", err)
	}
	partial["GapDuration"] = fieldGapDuration
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("projection: SessionWindow._marshalJSONSessionWindow: struct; %w", err)
	}
	return result, nil
}
func (r *SessionWindow) _marshalJSONtime_Duration(x time.Duration) ([]byte, error) {
	result, err := shared.JSONMarshal[time.Duration](x)
	if err != nil {
		return nil, fmt.Errorf("projection: SessionWindow._marshalJSONtime_Duration:; %w", err)
	}
	return result, nil
}
func (r *SessionWindow) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONSessionWindow(data)
	if err != nil {
		return fmt.Errorf("projection: SessionWindow.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *SessionWindow) _unmarshalJSONSessionWindow(data []byte) (SessionWindow, error) {
	result := SessionWindow{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("projection: SessionWindow._unmarshalJSONSessionWindow: native struct unwrap; %w", err)
	}
	if fieldGapDuration, ok := partial["GapDuration"]; ok {
		result.GapDuration, err = r._unmarshalJSONtime_Duration(fieldGapDuration)
		if err != nil {
			return result, fmt.Errorf("projection: SessionWindow._unmarshalJSONSessionWindow: field GapDuration; %w", err)
		}
	}
	return result, nil
}
func (r *SessionWindow) _unmarshalJSONtime_Duration(data []byte) (time.Duration, error) {
	result, err := shared.JSONUnmarshal[time.Duration](data)
	if err != nil {
		return result, fmt.Errorf("projection: SessionWindow._unmarshalJSONtime_Duration: native ref unwrap; %w", err)
	}
	return result, nil
}

func SlidingWindowFromJSON(x []byte) (*SlidingWindow, error) {
	result := new(SlidingWindow)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func SlidingWindowToJSON(x *SlidingWindow) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*SlidingWindow)(nil)
	_ json.Marshaler   = (*SlidingWindow)(nil)
)

func (r *SlidingWindow) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONSlidingWindow(*r)
}
func (r *SlidingWindow) _marshalJSONSlidingWindow(x SlidingWindow) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldWidth []byte
	fieldWidth, err = r._marshalJSONtime_Duration(x.Width)
	if err != nil {
		return nil, fmt.Errorf("projection: SlidingWindow._marshalJSONSlidingWindow: field name Width; %w", err)
	}
	partial["Width"] = fieldWidth
	var fieldPeriod []byte
	fieldPeriod, err = r._marshalJSONtime_Duration(x.Period)
	if err != nil {
		return nil, fmt.Errorf("projection: SlidingWindow._marshalJSONSlidingWindow: field name Period; %w", err)
	}
	partial["Period"] = fieldPeriod
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("projection: SlidingWindow._marshalJSONSlidingWindow: struct; %w", err)
	}
	return result, nil
}
func (r *SlidingWindow) _marshalJSONtime_Duration(x time.Duration) ([]byte, error) {
	result, err := shared.JSONMarshal[time.Duration](x)
	if err != nil {
		return nil, fmt.Errorf("projection: SlidingWindow._marshalJSONtime_Duration:; %w", err)
	}
	return result, nil
}
func (r *SlidingWindow) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONSlidingWindow(data)
	if err != nil {
		return fmt.Errorf("projection: SlidingWindow.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *SlidingWindow) _unmarshalJSONSlidingWindow(data []byte) (SlidingWindow, error) {
	result := SlidingWindow{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("projection: SlidingWindow._unmarshalJSONSlidingWindow: native struct unwrap; %w", err)
	}
	if fieldWidth, ok := partial["Width"]; ok {
		result.Width, err = r._unmarshalJSONtime_Duration(fieldWidth)
		if err != nil {
			return result, fmt.Errorf("projection: SlidingWindow._unmarshalJSONSlidingWindow: field Width; %w", err)
		}
	}
	if fieldPeriod, ok := partial["Period"]; ok {
		result.Period, err = r._unmarshalJSONtime_Duration(fieldPeriod)
		if err != nil {
			return result, fmt.Errorf("projection: SlidingWindow._unmarshalJSONSlidingWindow: field Period; %w", err)
		}
	}
	return result, nil
}
func (r *SlidingWindow) _unmarshalJSONtime_Duration(data []byte) (time.Duration, error) {
	result, err := shared.JSONUnmarshal[time.Duration](data)
	if err != nil {
		return result, fmt.Errorf("projection: SlidingWindow._unmarshalJSONtime_Duration: native ref unwrap; %w", err)
	}
	return result, nil
}

func FixedWindowFromJSON(x []byte) (*FixedWindow, error) {
	result := new(FixedWindow)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func FixedWindowToJSON(x *FixedWindow) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*FixedWindow)(nil)
	_ json.Marshaler   = (*FixedWindow)(nil)
)

func (r *FixedWindow) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONFixedWindow(*r)
}
func (r *FixedWindow) _marshalJSONFixedWindow(x FixedWindow) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldWidth []byte
	fieldWidth, err = r._marshalJSONtime_Duration(x.Width)
	if err != nil {
		return nil, fmt.Errorf("projection: FixedWindow._marshalJSONFixedWindow: field name Width; %w", err)
	}
	partial["Width"] = fieldWidth
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("projection: FixedWindow._marshalJSONFixedWindow: struct; %w", err)
	}
	return result, nil
}
func (r *FixedWindow) _marshalJSONtime_Duration(x time.Duration) ([]byte, error) {
	result, err := shared.JSONMarshal[time.Duration](x)
	if err != nil {
		return nil, fmt.Errorf("projection: FixedWindow._marshalJSONtime_Duration:; %w", err)
	}
	return result, nil
}
func (r *FixedWindow) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONFixedWindow(data)
	if err != nil {
		return fmt.Errorf("projection: FixedWindow.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *FixedWindow) _unmarshalJSONFixedWindow(data []byte) (FixedWindow, error) {
	result := FixedWindow{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("projection: FixedWindow._unmarshalJSONFixedWindow: native struct unwrap; %w", err)
	}
	if fieldWidth, ok := partial["Width"]; ok {
		result.Width, err = r._unmarshalJSONtime_Duration(fieldWidth)
		if err != nil {
			return result, fmt.Errorf("projection: FixedWindow._unmarshalJSONFixedWindow: field Width; %w", err)
		}
	}
	return result, nil
}
func (r *FixedWindow) _unmarshalJSONtime_Duration(data []byte) (time.Duration, error) {
	result, err := shared.JSONUnmarshal[time.Duration](data)
	if err != nil {
		return result, fmt.Errorf("projection: FixedWindow._unmarshalJSONtime_Duration: native ref unwrap; %w", err)
	}
	return result, nil
}

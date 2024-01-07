// Code generated by mkunion. DO NOT EDIT.
package projection

import (
	"encoding/json"
	"fmt"
	"github.com/widmogrod/mkunion/x/schema"
	"github.com/widmogrod/mkunion/x/shape"
	"github.com/widmogrod/mkunion/x/shared"
)

func init() {
	shape.Register(MessageShape())
	shape.Register(WindowShape())
	shape.Register(ItemTypeShape())
	shape.Register(ItemShape())
	shape.Register(StatsShape())
	shape.Register(EventTimeShape())
	shape.Register(ItemGroupedByWindowShape())
	shape.Register(ItemGroupedByKeyShape())
}

var (
	_ json.Unmarshaler = (*Message)(nil)
	_ json.Marshaler   = (*Message)(nil)
)

func (r *Message) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONMessage(*r)
}
func (r *Message) _marshalJSONMessage(x Message) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldOffset []byte
	fieldOffset, err = r._marshalJSONint(x.Offset)
	if err != nil {
		return nil, fmt.Errorf("projection: Message._marshalJSONMessage: field name Offset; %w", err)
	}
	partial["Offset"] = fieldOffset
	var fieldKey []byte
	fieldKey, err = r._marshalJSONstring(x.Key)
	if err != nil {
		return nil, fmt.Errorf("projection: Message._marshalJSONMessage: field name Key; %w", err)
	}
	partial["Key"] = fieldKey
	var fieldItem []byte
	fieldItem, err = r._marshalJSONPtrItem(x.Item)
	if err != nil {
		return nil, fmt.Errorf("projection: Message._marshalJSONMessage: field name Item; %w", err)
	}
	if fieldItem != nil {
		partial["Item"] = fieldItem
	}
	var fieldWatermark []byte
	fieldWatermark, err = r._marshalJSONPtrint64(x.Watermark)
	if err != nil {
		return nil, fmt.Errorf("projection: Message._marshalJSONMessage: field name Watermark; %w", err)
	}
	if fieldWatermark != nil {
		partial["Watermark"] = fieldWatermark
	}
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("projection: Message._marshalJSONMessage: struct; %w", err)
	}
	return result, nil
}
func (r *Message) _marshalJSONint(x int) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("projection: Message._marshalJSONint:; %w", err)
	}
	return result, nil
}
func (r *Message) _marshalJSONstring(x string) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("projection: Message._marshalJSONstring:; %w", err)
	}
	return result, nil
}
func (r *Message) _marshalJSONPtrItem(x *Item) ([]byte, error) {
	if x == nil {
		return nil, nil
	}
	return r._marshalJSONItem(*x)
}
func (r *Message) _marshalJSONItem(x Item) ([]byte, error) {
	result, err := shared.JSONMarshal[Item](x)
	if err != nil {
		return nil, fmt.Errorf("projection: Message._marshalJSONItem:; %w", err)
	}
	return result, nil
}
func (r *Message) _marshalJSONPtrint64(x *int64) ([]byte, error) {
	if x == nil {
		return nil, nil
	}
	return r._marshalJSONint64(*x)
}
func (r *Message) _marshalJSONint64(x int64) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("projection: Message._marshalJSONint64:; %w", err)
	}
	return result, nil
}
func (r *Message) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONMessage(data)
	if err != nil {
		return fmt.Errorf("projection: Message.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *Message) _unmarshalJSONMessage(data []byte) (Message, error) {
	result := Message{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("projection: Message._unmarshalJSONMessage: native struct unwrap; %w", err)
	}
	if fieldOffset, ok := partial["Offset"]; ok {
		result.Offset, err = r._unmarshalJSONint(fieldOffset)
		if err != nil {
			return result, fmt.Errorf("projection: Message._unmarshalJSONMessage: field Offset; %w", err)
		}
	}
	if fieldKey, ok := partial["Key"]; ok {
		result.Key, err = r._unmarshalJSONstring(fieldKey)
		if err != nil {
			return result, fmt.Errorf("projection: Message._unmarshalJSONMessage: field Key; %w", err)
		}
	}
	if fieldItem, ok := partial["Item"]; ok {
		result.Item, err = r._unmarshalJSONPtrItem(fieldItem)
		if err != nil {
			return result, fmt.Errorf("projection: Message._unmarshalJSONMessage: field Item; %w", err)
		}
	}
	if fieldWatermark, ok := partial["Watermark"]; ok {
		result.Watermark, err = r._unmarshalJSONPtrint64(fieldWatermark)
		if err != nil {
			return result, fmt.Errorf("projection: Message._unmarshalJSONMessage: field Watermark; %w", err)
		}
	}
	return result, nil
}
func (r *Message) _unmarshalJSONint(data []byte) (int, error) {
	var result int
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("projection: Message._unmarshalJSONint: native primitive unwrap; %w", err)
	}
	return result, nil
}
func (r *Message) _unmarshalJSONstring(data []byte) (string, error) {
	var result string
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("projection: Message._unmarshalJSONstring: native primitive unwrap; %w", err)
	}
	return result, nil
}
func (r *Message) _unmarshalJSONPtrItem(data []byte) (*Item, error) {
	if len(data) == 0 {
		return nil, nil
	}
	if string(data[:4]) == "null" {
		return nil, nil
	}
	result, err := r._unmarshalJSONItem(data)
	if err != nil {
		return nil, fmt.Errorf("projection: Message._unmarshalJSONPtrItem: pointer; %w", err)
	}
	return &result, nil
}
func (r *Message) _unmarshalJSONItem(data []byte) (Item, error) {
	result, err := shared.JSONUnmarshal[Item](data)
	if err != nil {
		return result, fmt.Errorf("projection: Message._unmarshalJSONItem: native ref unwrap; %w", err)
	}
	return result, nil
}
func (r *Message) _unmarshalJSONPtrint64(data []byte) (*int64, error) {
	if len(data) == 0 {
		return nil, nil
	}
	if string(data[:4]) == "null" {
		return nil, nil
	}
	result, err := r._unmarshalJSONint64(data)
	if err != nil {
		return nil, fmt.Errorf("projection: Message._unmarshalJSONPtrint64: pointer; %w", err)
	}
	return &result, nil
}
func (r *Message) _unmarshalJSONint64(data []byte) (int64, error) {
	var result int64
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("projection: Message._unmarshalJSONint64: native primitive unwrap; %w", err)
	}
	return result, nil
}
func MessageShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Message",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless/projection",
		Fields: []*shape.FieldLike{
			{
				Name: "Offset",
				Type: &shape.PrimitiveLike{
					Kind: &shape.NumberLike{
						Kind: &shape.Int{},
					},
				},
			},
			{
				Name: "Key",
				Type: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
			},
			{
				Name: "Item",
				Type: &shape.PointerLike{
					Type: &shape.RefName{
						Name:          "Item",
						PkgName:       "projection",
						PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless/projection",
					},
				},
			},
			{
				Name: "Watermark",
				Type: &shape.PointerLike{
					Type: &shape.PrimitiveLike{
						Kind: &shape.NumberLike{
							Kind: &shape.Int64{},
						},
					},
				},
			},
		},
		Tags: map[string]shape.Tag{
			"serde": {
				Value: "json",
			},
		},
	}
}

var (
	_ json.Unmarshaler = (*Window)(nil)
	_ json.Marshaler   = (*Window)(nil)
)

func (r *Window) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONWindow(*r)
}
func (r *Window) _marshalJSONWindow(x Window) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldStart []byte
	fieldStart, err = r._marshalJSONint64(x.Start)
	if err != nil {
		return nil, fmt.Errorf("projection: Window._marshalJSONWindow: field name Start; %w", err)
	}
	partial["Start"] = fieldStart
	var fieldEnd []byte
	fieldEnd, err = r._marshalJSONint64(x.End)
	if err != nil {
		return nil, fmt.Errorf("projection: Window._marshalJSONWindow: field name End; %w", err)
	}
	partial["End"] = fieldEnd
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("projection: Window._marshalJSONWindow: struct; %w", err)
	}
	return result, nil
}
func (r *Window) _marshalJSONint64(x int64) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("projection: Window._marshalJSONint64:; %w", err)
	}
	return result, nil
}
func (r *Window) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONWindow(data)
	if err != nil {
		return fmt.Errorf("projection: Window.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *Window) _unmarshalJSONWindow(data []byte) (Window, error) {
	result := Window{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("projection: Window._unmarshalJSONWindow: native struct unwrap; %w", err)
	}
	if fieldStart, ok := partial["Start"]; ok {
		result.Start, err = r._unmarshalJSONint64(fieldStart)
		if err != nil {
			return result, fmt.Errorf("projection: Window._unmarshalJSONWindow: field Start; %w", err)
		}
	}
	if fieldEnd, ok := partial["End"]; ok {
		result.End, err = r._unmarshalJSONint64(fieldEnd)
		if err != nil {
			return result, fmt.Errorf("projection: Window._unmarshalJSONWindow: field End; %w", err)
		}
	}
	return result, nil
}
func (r *Window) _unmarshalJSONint64(data []byte) (int64, error) {
	var result int64
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("projection: Window._unmarshalJSONint64: native primitive unwrap; %w", err)
	}
	return result, nil
}
func WindowShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Window",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless/projection",
		Fields: []*shape.FieldLike{
			{
				Name: "Start",
				Type: &shape.PrimitiveLike{
					Kind: &shape.NumberLike{
						Kind: &shape.Int64{},
					},
				},
			},
			{
				Name: "End",
				Type: &shape.PrimitiveLike{
					Kind: &shape.NumberLike{
						Kind: &shape.Int64{},
					},
				},
			},
		},
		Tags: map[string]shape.Tag{
			"serde": {
				Value: "json",
			},
		},
	}
}

var (
	_ json.Unmarshaler = (*ItemType)(nil)
	_ json.Marshaler   = (*ItemType)(nil)
)

func (r *ItemType) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONItemType(*r)
}
func (r *ItemType) _marshalJSONItemType(x ItemType) ([]byte, error) {
	return r._marshalJSONuint8(uint8(x))
}
func (r *ItemType) _marshalJSONuint8(x uint8) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("projection: ItemType._marshalJSONuint8:; %w", err)
	}
	return result, nil
}
func (r *ItemType) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONItemType(data)
	if err != nil {
		return fmt.Errorf("projection: ItemType.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *ItemType) _unmarshalJSONItemType(data []byte) (ItemType, error) {
	var result ItemType
	intermidiary, err := r._unmarshalJSONuint8(data)
	if err != nil {
		return result, fmt.Errorf("projection: ItemType._unmarshalJSONItemType: alias; %w", err)
	}
	result = ItemType(intermidiary)
	return result, nil
}
func (r *ItemType) _unmarshalJSONuint8(data []byte) (uint8, error) {
	var result uint8
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("projection: ItemType._unmarshalJSONuint8: native primitive unwrap; %w", err)
	}
	return result, nil
}
func ItemTypeShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "ItemType",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless/projection",
		Tags: map[string]shape.Tag{
			"serde": {
				Value: "json",
			},
		},
		Type: &shape.PrimitiveLike{
			Kind: &shape.NumberLike{
				Kind: &shape.UInt8{},
			},
		},
	}
}

var (
	_ json.Unmarshaler = (*Item)(nil)
	_ json.Marshaler   = (*Item)(nil)
)

func (r *Item) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONItem(*r)
}
func (r *Item) _marshalJSONItem(x Item) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldKey []byte
	fieldKey, err = r._marshalJSONstring(x.Key)
	if err != nil {
		return nil, fmt.Errorf("projection: Item._marshalJSONItem: field name Key; %w", err)
	}
	partial["Key"] = fieldKey
	var fieldData []byte
	fieldData, err = r._marshalJSONschema_Schema(x.Data)
	if err != nil {
		return nil, fmt.Errorf("projection: Item._marshalJSONItem: field name Data; %w", err)
	}
	partial["Data"] = fieldData
	var fieldEventTime []byte
	fieldEventTime, err = r._marshalJSONEventTime(x.EventTime)
	if err != nil {
		return nil, fmt.Errorf("projection: Item._marshalJSONItem: field name EventTime; %w", err)
	}
	partial["EventTime"] = fieldEventTime
	var fieldWindow []byte
	fieldWindow, err = r._marshalJSONPtrWindow(x.Window)
	if err != nil {
		return nil, fmt.Errorf("projection: Item._marshalJSONItem: field name Window; %w", err)
	}
	if fieldWindow != nil {
		partial["Window"] = fieldWindow
	}
	var fieldType []byte
	fieldType, err = r._marshalJSONItemType(x.Type)
	if err != nil {
		return nil, fmt.Errorf("projection: Item._marshalJSONItem: field name Type; %w", err)
	}
	partial["Type"] = fieldType
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("projection: Item._marshalJSONItem: struct; %w", err)
	}
	return result, nil
}
func (r *Item) _marshalJSONstring(x string) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("projection: Item._marshalJSONstring:; %w", err)
	}
	return result, nil
}
func (r *Item) _marshalJSONschema_Schema(x schema.Schema) ([]byte, error) {
	result, err := shared.JSONMarshal[schema.Schema](x)
	if err != nil {
		return nil, fmt.Errorf("projection: Item._marshalJSONschema_Schema:; %w", err)
	}
	return result, nil
}
func (r *Item) _marshalJSONEventTime(x EventTime) ([]byte, error) {
	result, err := shared.JSONMarshal[EventTime](x)
	if err != nil {
		return nil, fmt.Errorf("projection: Item._marshalJSONEventTime:; %w", err)
	}
	return result, nil
}
func (r *Item) _marshalJSONPtrWindow(x *Window) ([]byte, error) {
	if x == nil {
		return nil, nil
	}
	return r._marshalJSONWindow(*x)
}
func (r *Item) _marshalJSONWindow(x Window) ([]byte, error) {
	result, err := shared.JSONMarshal[Window](x)
	if err != nil {
		return nil, fmt.Errorf("projection: Item._marshalJSONWindow:; %w", err)
	}
	return result, nil
}
func (r *Item) _marshalJSONItemType(x ItemType) ([]byte, error) {
	result, err := shared.JSONMarshal[ItemType](x)
	if err != nil {
		return nil, fmt.Errorf("projection: Item._marshalJSONItemType:; %w", err)
	}
	return result, nil
}
func (r *Item) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONItem(data)
	if err != nil {
		return fmt.Errorf("projection: Item.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *Item) _unmarshalJSONItem(data []byte) (Item, error) {
	result := Item{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("projection: Item._unmarshalJSONItem: native struct unwrap; %w", err)
	}
	if fieldKey, ok := partial["Key"]; ok {
		result.Key, err = r._unmarshalJSONstring(fieldKey)
		if err != nil {
			return result, fmt.Errorf("projection: Item._unmarshalJSONItem: field Key; %w", err)
		}
	}
	if fieldData, ok := partial["Data"]; ok {
		result.Data, err = r._unmarshalJSONschema_Schema(fieldData)
		if err != nil {
			return result, fmt.Errorf("projection: Item._unmarshalJSONItem: field Data; %w", err)
		}
	}
	if fieldEventTime, ok := partial["EventTime"]; ok {
		result.EventTime, err = r._unmarshalJSONEventTime(fieldEventTime)
		if err != nil {
			return result, fmt.Errorf("projection: Item._unmarshalJSONItem: field EventTime; %w", err)
		}
	}
	if fieldWindow, ok := partial["Window"]; ok {
		result.Window, err = r._unmarshalJSONPtrWindow(fieldWindow)
		if err != nil {
			return result, fmt.Errorf("projection: Item._unmarshalJSONItem: field Window; %w", err)
		}
	}
	if fieldType, ok := partial["Type"]; ok {
		result.Type, err = r._unmarshalJSONItemType(fieldType)
		if err != nil {
			return result, fmt.Errorf("projection: Item._unmarshalJSONItem: field Type; %w", err)
		}
	}
	return result, nil
}
func (r *Item) _unmarshalJSONstring(data []byte) (string, error) {
	var result string
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("projection: Item._unmarshalJSONstring: native primitive unwrap; %w", err)
	}
	return result, nil
}
func (r *Item) _unmarshalJSONschema_Schema(data []byte) (schema.Schema, error) {
	result, err := shared.JSONUnmarshal[schema.Schema](data)
	if err != nil {
		return result, fmt.Errorf("projection: Item._unmarshalJSONschema_Schema: native ref unwrap; %w", err)
	}
	return result, nil
}
func (r *Item) _unmarshalJSONEventTime(data []byte) (EventTime, error) {
	result, err := shared.JSONUnmarshal[EventTime](data)
	if err != nil {
		return result, fmt.Errorf("projection: Item._unmarshalJSONEventTime: native ref unwrap; %w", err)
	}
	return result, nil
}
func (r *Item) _unmarshalJSONPtrWindow(data []byte) (*Window, error) {
	if len(data) == 0 {
		return nil, nil
	}
	if string(data[:4]) == "null" {
		return nil, nil
	}
	result, err := r._unmarshalJSONWindow(data)
	if err != nil {
		return nil, fmt.Errorf("projection: Item._unmarshalJSONPtrWindow: pointer; %w", err)
	}
	return &result, nil
}
func (r *Item) _unmarshalJSONWindow(data []byte) (Window, error) {
	result, err := shared.JSONUnmarshal[Window](data)
	if err != nil {
		return result, fmt.Errorf("projection: Item._unmarshalJSONWindow: native ref unwrap; %w", err)
	}
	return result, nil
}
func (r *Item) _unmarshalJSONItemType(data []byte) (ItemType, error) {
	result, err := shared.JSONUnmarshal[ItemType](data)
	if err != nil {
		return result, fmt.Errorf("projection: Item._unmarshalJSONItemType: native ref unwrap; %w", err)
	}
	return result, nil
}
func ItemShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Item",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless/projection",
		Fields: []*shape.FieldLike{
			{
				Name: "Key",
				Type: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
			},
			{
				Name: "Data",
				Type: &shape.RefName{
					Name:          "Schema",
					PkgName:       "schema",
					PkgImportName: "github.com/widmogrod/mkunion/x/schema",
				},
			},
			{
				Name: "EventTime",
				Type: &shape.RefName{
					Name:          "EventTime",
					PkgName:       "projection",
					PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless/projection",
				},
			},
			{
				Name: "Window",
				Type: &shape.PointerLike{
					Type: &shape.RefName{
						Name:          "Window",
						PkgName:       "projection",
						PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless/projection",
					},
				},
			},
			{
				Name: "Type",
				Type: &shape.RefName{
					Name:          "ItemType",
					PkgName:       "projection",
					PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless/projection",
				},
			},
		},
		Tags: map[string]shape.Tag{
			"serde": {
				Value: "json",
			},
		},
	}
}

func StatsShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "Stats",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless/projection",
		Tags: map[string]shape.Tag{
			"serde": {
				Value: "json",
			},
		},
		IsAlias: true,
		Type: &shape.MapLike{
			Key: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
			Val: &shape.PrimitiveLike{
				Kind: &shape.NumberLike{
					Kind: &shape.Int{},
				},
			},
		},
	}
}

func EventTimeShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "EventTime",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless/projection",
		Tags: map[string]shape.Tag{
			"serde": {
				Value: "json",
			},
		},
		IsAlias: true,
		Type: &shape.PrimitiveLike{
			Kind: &shape.NumberLike{
				Kind: &shape.Int64{},
			},
		},
	}
}

var (
	_ json.Unmarshaler = (*ItemGroupedByWindow)(nil)
	_ json.Marshaler   = (*ItemGroupedByWindow)(nil)
)

func (r *ItemGroupedByWindow) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONItemGroupedByWindow(*r)
}
func (r *ItemGroupedByWindow) _marshalJSONItemGroupedByWindow(x ItemGroupedByWindow) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldKey []byte
	fieldKey, err = r._marshalJSONstring(x.Key)
	if err != nil {
		return nil, fmt.Errorf("projection: ItemGroupedByWindow._marshalJSONItemGroupedByWindow: field name Key; %w", err)
	}
	partial["Key"] = fieldKey
	var fieldData []byte
	fieldData, err = r._marshalJSONPtrschema_List(x.Data)
	if err != nil {
		return nil, fmt.Errorf("projection: ItemGroupedByWindow._marshalJSONItemGroupedByWindow: field name Data; %w", err)
	}
	if fieldData != nil {
		partial["Data"] = fieldData
	}
	var fieldWindow []byte
	fieldWindow, err = r._marshalJSONPtrWindow(x.Window)
	if err != nil {
		return nil, fmt.Errorf("projection: ItemGroupedByWindow._marshalJSONItemGroupedByWindow: field name Window; %w", err)
	}
	if fieldWindow != nil {
		partial["Window"] = fieldWindow
	}
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("projection: ItemGroupedByWindow._marshalJSONItemGroupedByWindow: struct; %w", err)
	}
	return result, nil
}
func (r *ItemGroupedByWindow) _marshalJSONstring(x string) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("projection: ItemGroupedByWindow._marshalJSONstring:; %w", err)
	}
	return result, nil
}
func (r *ItemGroupedByWindow) _marshalJSONPtrschema_List(x *schema.List) ([]byte, error) {
	if x == nil {
		return nil, nil
	}
	return r._marshalJSONschema_List(*x)
}
func (r *ItemGroupedByWindow) _marshalJSONschema_List(x schema.List) ([]byte, error) {
	result, err := shared.JSONMarshal[schema.List](x)
	if err != nil {
		return nil, fmt.Errorf("projection: ItemGroupedByWindow._marshalJSONschema_List:; %w", err)
	}
	return result, nil
}
func (r *ItemGroupedByWindow) _marshalJSONPtrWindow(x *Window) ([]byte, error) {
	if x == nil {
		return nil, nil
	}
	return r._marshalJSONWindow(*x)
}
func (r *ItemGroupedByWindow) _marshalJSONWindow(x Window) ([]byte, error) {
	result, err := shared.JSONMarshal[Window](x)
	if err != nil {
		return nil, fmt.Errorf("projection: ItemGroupedByWindow._marshalJSONWindow:; %w", err)
	}
	return result, nil
}
func (r *ItemGroupedByWindow) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONItemGroupedByWindow(data)
	if err != nil {
		return fmt.Errorf("projection: ItemGroupedByWindow.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *ItemGroupedByWindow) _unmarshalJSONItemGroupedByWindow(data []byte) (ItemGroupedByWindow, error) {
	result := ItemGroupedByWindow{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("projection: ItemGroupedByWindow._unmarshalJSONItemGroupedByWindow: native struct unwrap; %w", err)
	}
	if fieldKey, ok := partial["Key"]; ok {
		result.Key, err = r._unmarshalJSONstring(fieldKey)
		if err != nil {
			return result, fmt.Errorf("projection: ItemGroupedByWindow._unmarshalJSONItemGroupedByWindow: field Key; %w", err)
		}
	}
	if fieldData, ok := partial["Data"]; ok {
		result.Data, err = r._unmarshalJSONPtrschema_List(fieldData)
		if err != nil {
			return result, fmt.Errorf("projection: ItemGroupedByWindow._unmarshalJSONItemGroupedByWindow: field Data; %w", err)
		}
	}
	if fieldWindow, ok := partial["Window"]; ok {
		result.Window, err = r._unmarshalJSONPtrWindow(fieldWindow)
		if err != nil {
			return result, fmt.Errorf("projection: ItemGroupedByWindow._unmarshalJSONItemGroupedByWindow: field Window; %w", err)
		}
	}
	return result, nil
}
func (r *ItemGroupedByWindow) _unmarshalJSONstring(data []byte) (string, error) {
	var result string
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("projection: ItemGroupedByWindow._unmarshalJSONstring: native primitive unwrap; %w", err)
	}
	return result, nil
}
func (r *ItemGroupedByWindow) _unmarshalJSONPtrschema_List(data []byte) (*schema.List, error) {
	if len(data) == 0 {
		return nil, nil
	}
	if string(data[:4]) == "null" {
		return nil, nil
	}
	result, err := r._unmarshalJSONschema_List(data)
	if err != nil {
		return nil, fmt.Errorf("projection: ItemGroupedByWindow._unmarshalJSONPtrschema_List: pointer; %w", err)
	}
	return &result, nil
}
func (r *ItemGroupedByWindow) _unmarshalJSONschema_List(data []byte) (schema.List, error) {
	result, err := shared.JSONUnmarshal[schema.List](data)
	if err != nil {
		return result, fmt.Errorf("projection: ItemGroupedByWindow._unmarshalJSONschema_List: native ref unwrap; %w", err)
	}
	return result, nil
}
func (r *ItemGroupedByWindow) _unmarshalJSONPtrWindow(data []byte) (*Window, error) {
	if len(data) == 0 {
		return nil, nil
	}
	if string(data[:4]) == "null" {
		return nil, nil
	}
	result, err := r._unmarshalJSONWindow(data)
	if err != nil {
		return nil, fmt.Errorf("projection: ItemGroupedByWindow._unmarshalJSONPtrWindow: pointer; %w", err)
	}
	return &result, nil
}
func (r *ItemGroupedByWindow) _unmarshalJSONWindow(data []byte) (Window, error) {
	result, err := shared.JSONUnmarshal[Window](data)
	if err != nil {
		return result, fmt.Errorf("projection: ItemGroupedByWindow._unmarshalJSONWindow: native ref unwrap; %w", err)
	}
	return result, nil
}
func ItemGroupedByWindowShape() shape.Shape {
	return &shape.StructLike{
		Name:          "ItemGroupedByWindow",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless/projection",
		Fields: []*shape.FieldLike{
			{
				Name: "Key",
				Type: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
			},
			{
				Name: "Data",
				Type: &shape.PointerLike{
					Type: &shape.RefName{
						Name:          "List",
						PkgName:       "schema",
						PkgImportName: "github.com/widmogrod/mkunion/x/schema",
					},
				},
			},
			{
				Name: "Window",
				Type: &shape.PointerLike{
					Type: &shape.RefName{
						Name:          "Window",
						PkgName:       "projection",
						PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless/projection",
					},
				},
			},
		},
		Tags: map[string]shape.Tag{
			"serde": {
				Value: "json",
			},
		},
	}
}

var (
	_ json.Unmarshaler = (*ItemGroupedByKey)(nil)
	_ json.Marshaler   = (*ItemGroupedByKey)(nil)
)

func (r *ItemGroupedByKey) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONItemGroupedByKey(*r)
}
func (r *ItemGroupedByKey) _marshalJSONItemGroupedByKey(x ItemGroupedByKey) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldKey []byte
	fieldKey, err = r._marshalJSONstring(x.Key)
	if err != nil {
		return nil, fmt.Errorf("projection: ItemGroupedByKey._marshalJSONItemGroupedByKey: field name Key; %w", err)
	}
	partial["Key"] = fieldKey
	var fieldData []byte
	fieldData, err = r._marshalJSONSliceItem(x.Data)
	if err != nil {
		return nil, fmt.Errorf("projection: ItemGroupedByKey._marshalJSONItemGroupedByKey: field name Data; %w", err)
	}
	partial["Data"] = fieldData
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("projection: ItemGroupedByKey._marshalJSONItemGroupedByKey: struct; %w", err)
	}
	return result, nil
}
func (r *ItemGroupedByKey) _marshalJSONstring(x string) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("projection: ItemGroupedByKey._marshalJSONstring:; %w", err)
	}
	return result, nil
}
func (r *ItemGroupedByKey) _marshalJSONSliceItem(x []Item) ([]byte, error) {
	partial := make([]json.RawMessage, len(x))
	for i, v := range x {
		item, err := r._marshalJSONItem(v)
		if err != nil {
			return nil, fmt.Errorf("projection: ItemGroupedByKey._marshalJSONSliceItem: at index %d; %w", i, err)
		}
		partial[i] = item
	}
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("projection: ItemGroupedByKey._marshalJSONSliceItem:; %w", err)
	}
	return result, nil
}
func (r *ItemGroupedByKey) _marshalJSONItem(x Item) ([]byte, error) {
	result, err := shared.JSONMarshal[Item](x)
	if err != nil {
		return nil, fmt.Errorf("projection: ItemGroupedByKey._marshalJSONItem:; %w", err)
	}
	return result, nil
}
func (r *ItemGroupedByKey) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONItemGroupedByKey(data)
	if err != nil {
		return fmt.Errorf("projection: ItemGroupedByKey.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *ItemGroupedByKey) _unmarshalJSONItemGroupedByKey(data []byte) (ItemGroupedByKey, error) {
	result := ItemGroupedByKey{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("projection: ItemGroupedByKey._unmarshalJSONItemGroupedByKey: native struct unwrap; %w", err)
	}
	if fieldKey, ok := partial["Key"]; ok {
		result.Key, err = r._unmarshalJSONstring(fieldKey)
		if err != nil {
			return result, fmt.Errorf("projection: ItemGroupedByKey._unmarshalJSONItemGroupedByKey: field Key; %w", err)
		}
	}
	if fieldData, ok := partial["Data"]; ok {
		result.Data, err = r._unmarshalJSONSliceItem(fieldData)
		if err != nil {
			return result, fmt.Errorf("projection: ItemGroupedByKey._unmarshalJSONItemGroupedByKey: field Data; %w", err)
		}
	}
	return result, nil
}
func (r *ItemGroupedByKey) _unmarshalJSONstring(data []byte) (string, error) {
	var result string
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("projection: ItemGroupedByKey._unmarshalJSONstring: native primitive unwrap; %w", err)
	}
	return result, nil
}
func (r *ItemGroupedByKey) _unmarshalJSONSliceItem(data []byte) ([]Item, error) {
	result := make([]Item, 0)
	var partial []json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("projection: ItemGroupedByKey._unmarshalJSONSliceItem: native list unwrap; %w", err)
	}
	for i, v := range partial {
		item, err := r._unmarshalJSONItem(v)
		if err != nil {
			return result, fmt.Errorf("projection: ItemGroupedByKey._unmarshalJSONSliceItem: at index %d; %w", i, err)
		}
		result = append(result, item)
	}
	return result, nil
}
func (r *ItemGroupedByKey) _unmarshalJSONItem(data []byte) (Item, error) {
	result, err := shared.JSONUnmarshal[Item](data)
	if err != nil {
		return result, fmt.Errorf("projection: ItemGroupedByKey._unmarshalJSONItem: native ref unwrap; %w", err)
	}
	return result, nil
}
func ItemGroupedByKeyShape() shape.Shape {
	return &shape.StructLike{
		Name:          "ItemGroupedByKey",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless/projection",
		Fields: []*shape.FieldLike{
			{
				Name: "Key",
				Type: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
			},
			{
				Name: "Data",
				Type: &shape.ListLike{
					Element: &shape.RefName{
						Name:          "Item",
						PkgName:       "projection",
						PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless/projection",
					},
				},
			},
		},
		Tags: map[string]shape.Tag{
			"serde": {
				Value: "json",
			},
		},
	}
}

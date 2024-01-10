// Code generated by mkunion. DO NOT EDIT.
package projection

import (
	"encoding/json"
	"fmt"
	"github.com/widmogrod/mkunion/x/schema"
	"github.com/widmogrod/mkunion/x/shared"
)

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
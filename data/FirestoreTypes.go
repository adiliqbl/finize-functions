package data

import (
	"encoding/json"
	"finize-functions.app/util"
	"time"
)

type BooleanValue struct {
	Value *bool `json:"booleanValue,omitempty"`
}

func (it BooleanValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(util.ValueOrNull(it.Value))
}

type IntValue struct {
	Value *int `json:"integerValue,omitempty"`
}

func (it IntValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(util.ValueOrNull(it.Value))
}

type DoubleValue struct {
	Value *float64 `json:"doubleValue,omitempty"`
}

func (it DoubleValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(util.ValueOrNull(it.Value))
}

type StringValue struct {
	Value *string `json:"stringValue,omitempty"`
}

func (it StringValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(util.ValueOrNull(it.Value))
}

type TimestampValue struct {
	Value *time.Time `json:"timestampValue"`
}

func (it TimestampValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(util.ValueOrNull(it.Value))
}

type ArrayValue[T any] struct {
	Value *[]T `json:"arrayValue,omitempty"`
}

func (it ArrayValue[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(util.ValueOrNull(it.Value))
}

type MapValue[T any] struct {
	Value *T `json:"mapValue,omitempty"`
}

func (it MapValue[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(util.ValueOrNull(it.Value))
}

type ReferenceValue struct {
	Reference *string `json:"referenceValue,omitempty"`
}

func (it ReferenceValue) Get() *string {
	if util.NullOrEmpty(it.Reference) {
		return nil
	} else {
		return it.Reference
	}
}

func (it ReferenceValue) MarshalJSON() ([]byte, error) {
	if util.NullOrEmpty(it.Reference) {
		return json.Marshal(nil)
	} else {
		return json.Marshal(&it.Reference)
	}
}

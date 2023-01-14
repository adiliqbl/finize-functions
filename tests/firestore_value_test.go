package tests

import (
	"encoding/json"
	"finize-functions.app/data"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFirebaseBooleanValue(t *testing.T) {
	validate[bool](t, true, marshal(data.BooleanValue{Value: util.Pointer(true)}))
	validateNil[bool](t, marshal(data.BooleanValue{Value: nil}))
}

func TestFirebaseIntValue(t *testing.T) {
	validate[int](t, 50, marshal(data.IntValue{Value: util.Pointer(50)}))
	validateNil[int](t, marshal(data.IntValue{Value: nil}))
}

func TestFirebaseDoubleValue(t *testing.T) {
	validate[float64](t, 50.0, marshal(data.DoubleValue{Value: util.Pointer(50.0)}))
	validateNil[float64](t, marshal(data.DoubleValue{Value: nil}))
}

func TestFirebaseStringValue(t *testing.T) {
	validate[string](t, "test", marshal(data.StringValue{Value: util.Pointer("test")}))
	validateNil[string](t, marshal(data.StringValue{Value: nil}))
}

func TestFirebaseTimestampValue(t *testing.T) {
	want, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z07:00")
	got := marshal(data.TimestampValue{Value: util.Pointer(want)})

	var value time.Time
	err := json.Unmarshal(got, &value)

	assert.Nil(t, err)
	assert.True(t, value.Equal(want))
}

func TestFirebaseArrayValue(t *testing.T) {
	validate[[]string](t, []string{"one", "two"}, marshal(data.ArrayValue[string]{Value: util.Pointer([]string{"one", "two"})}))
	validateNil[[]string](t, marshal(data.ArrayValue[string]{Value: nil}))
}

func TestFirebaseMapValue(t *testing.T) {
	want := map[string]interface{}{"key": "value"}
	validate[map[string]interface{}](t, want, marshal(data.MapValue[map[string]interface{}]{Value: util.Pointer(want)}))
	validateNil[map[string]interface{}](t, marshal(data.MapValue[map[string]interface{}]{Value: nil}))
}

func TestFirebaseReferenceValue(t *testing.T) {
	validate[string](t, "ref", marshal(data.ReferenceValue{Reference: util.Pointer("ref")}))
	validateNil[string](t, marshal(data.ReferenceValue{Reference: nil}))
}

func marshal(value any) []byte {
	marshaled, _ := json.Marshal(value)
	return marshaled
}

func validate[T any](t *testing.T, want any, marshal []byte) {
	var value T
	err := json.Unmarshal(marshal, &value)

	assert.Nil(t, err)
	assert.Equal(t, want, value)
}

func validateNil[T any](t *testing.T, marshal []byte) {
	var value *T
	err := json.Unmarshal(marshal, &value)

	assert.Nil(t, err)
	assert.Nil(t, value)
}

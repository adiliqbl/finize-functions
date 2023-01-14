package util

import "reflect"

func removeNulls(m map[string]interface{}) {
	val := reflect.ValueOf(m)
	for _, e := range val.MapKeys() {
		v := val.MapIndex(e)
		if v.IsNil() {
			delete(m, e.String())
			continue
		}
		switch t := v.Interface().(type) {
		case map[string]interface{}:
			removeNulls(t)
		}
	}
}

func Pointer[T any](value T) *T {
	return &value
}

func NullOrEmpty(str *string) bool {
	return str == nil || *str == ""
}

func ValueOrNull[T any](ptr *T) interface{} {
	if ptr == nil {
		return nil
	} else {
		return *ptr
	}
}

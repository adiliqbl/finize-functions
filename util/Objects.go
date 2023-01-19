package util

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

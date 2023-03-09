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

func GenerateKeywords(text string) []string {
	var keywords []string
	for i := 0; i < len(text); i++ {
		for j := i + 1; j <= len(text); j++ {
			keywords = append(keywords, text[i:j])
		}
	}
	return keywords
}

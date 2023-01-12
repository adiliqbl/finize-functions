package data

type FirestoreEvent[T any] struct {
	OldValue   FirestoreValue[T] `json:"oldValue"`
	Value      FirestoreValue[T] `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

package data

type FirestoreEvent[T any] struct {
	OldValue   FirestoreValue[T] `json:"oldValue"`
	Value      FirestoreValue[T] `json:"value"`
	UpdateMask struct {
		Fields []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

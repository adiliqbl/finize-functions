package data

import "strings"

type FirestoreEvent[T any] struct {
	OldValue   FirestoreValue[T] `json:"oldValue"`
	Value      FirestoreValue[T] `json:"value"`
	UpdateMask struct {
		Fields []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

func (event FirestoreEvent[T]) UserID() string {
	fullPath := strings.Split(event.Value.Name, "/documents/")[1]
	pathParts := strings.Split(fullPath, "/")
	return pathParts[1]
}

func (event FirestoreEvent[T]) Path() (path string, collection string, doc string) {
	fullPath := strings.Split(event.Value.Name, "/documents/")[1]
	pathParts := strings.Split(fullPath, "/")
	collection = pathParts[0]
	doc = strings.Join(pathParts[1:], "/")
	path = collection + "/" + doc
	return path, collection, doc
}

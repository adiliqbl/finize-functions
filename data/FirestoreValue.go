package data

import "time"

type FirestoreValue[T any] struct {
	CreateTime time.Time `json:"createTime"`
	Data       T         `json:"fields"`
	Name       string    `json:"name"`
	UpdateTime time.Time `json:"updateTime"`
}

package data

import "time"

type AuthEvent struct {
	UID      string `json:"uid"`
	Email    string `json:"email"`
	Metadata struct {
		CreatedAt time.Time `json:"createdAt"`
	} `json:"metadata"`
}

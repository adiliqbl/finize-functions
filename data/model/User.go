package model

import . "finize-functions.app/data"

type UserEvent struct {
	ID      StringValue            `json:"id"`
	Name    StringValue            `json:"name"`
	Email   StringValue            `json:"email"`
	Profile MapValue[ProfileEvent] `json:"profile"`
}

type User struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Profile Profile `json:"profile"`
}

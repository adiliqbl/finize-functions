package model

import . "finize-functions.app/data"

type User struct {
	ID    StringValue `json:"id"`
	Name  StringValue `json:"name"`
	Email StringValue `json:"email"`
	Image StringValue `json:"image,omitempty"`
}

package model

import . "finize-functions.app/data"

type ProfileEvent struct {
	Currency StringValue `json:"currency"`
}

type Profile struct {
	Currency string `json:"currency"`
}

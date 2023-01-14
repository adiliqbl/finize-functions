package service

import "finize-functions.app/data/model"

type UserService struct {
	db Firestore[model.UserEvent]
}

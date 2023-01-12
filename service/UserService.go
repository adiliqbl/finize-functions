package service

import "finize-functions/data/model"

type UserService struct {
	db Firestore[model.User]
}

package service

import (
	"context"
	"finize-functions.app/tests/fake"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	factory := fake.NewServiceFactory(context.Background(), "user")
	service := factory.UserService()

	id, err := service.Create(fake.NewUser(fake.NewUserEvent("user", "name", "email@test.com")))

	assert.Nil(t, err)
	assert.True(t, !util.NullOrEmpty(&id))
}

func TestGet(t *testing.T) {
	factory := fake.NewServiceFactory(context.Background(), "user")
	service := factory.UserService()

	testUser := fake.NewUser(fake.NewUserEvent("", "name", "email@test.com"))
	testUser.ID, _ = service.Create(testUser)

	user, _ := service.FindByID(testUser.ID)
	assert.Equal(t, testUser, *user)
}

func TestUpdate(t *testing.T) {
	factory := fake.NewServiceFactory(context.Background(), "user")
	service := factory.UserService()

	testUser := fake.NewUser(fake.NewUserEvent("", "name", "email@test.com"))
	testUser.ID, _ = service.Create(testUser)
	testUser.Name = "new name"

	doc, _ := util.MapTo[map[string]interface{}](testUser)
	ok, _ := service.Update(testUser.ID, doc)
	assert.True(t, ok)

	user, _ := service.FindByID(testUser.ID)
	assert.Equal(t, testUser, *user)
}

func TestDelete(t *testing.T) {
	factory := fake.NewServiceFactory(context.Background(), "user")
	service := factory.UserService()

	id, _ := service.Create(fake.NewUser(fake.NewUserEvent("", "name", "email@test.com")))
	ok, _ := service.Delete(id)
	assert.True(t, ok)

	user, _ := service.FindByID(id)
	assert.Nil(t, user)
}

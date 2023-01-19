package service

import (
	"finize-functions.app/tests/fake"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	id, err := userService.Create(fake.NewUser(fake.NewUserEvent("user", "name", "email@test.com")))

	assert.Nil(t, err)
	assert.True(t, !util.NullOrEmpty(&id))
}

func TestGet(t *testing.T) {
	testUser := fake.NewUser(fake.NewUserEvent("", "name", "email@test.com"))
	testUser.ID, _ = userService.Create(testUser)

	user, _ := userService.FindByID(testUser.ID)
	assert.Equal(t, testUser, *user)
}

func TestUpdate(t *testing.T) {
	testUser := fake.NewUser(fake.NewUserEvent("", "name", "email@test.com"))
	testUser.ID, _ = userService.Create(testUser)
	testUser.Name = "new name"

	doc, _ := util.MapTo[map[string]interface{}](testUser)
	ok, _ := userService.Update(testUser.ID, doc)
	assert.True(t, ok)

	user, _ := userService.FindByID(testUser.ID)
	assert.Equal(t, testUser, *user)
}

func TestDelete(t *testing.T) {
	id, _ := userService.Create(fake.NewUser(fake.NewUserEvent("", "name", "email@test.com")))
	ok, _ := userService.Delete(id)
	assert.True(t, ok)

	user, _ := userService.FindByID(id)
	assert.Nil(t, user)
}

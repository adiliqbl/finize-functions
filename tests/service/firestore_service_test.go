package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/tests/fake"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	id, err := userService.Create(fake.NewUser(fake.NewUserEvent("", "name", "email@test.com")))

	assert.Nil(t, err)
	assert.True(t, !util.NullOrEmpty(id))
}

func TestGet(t *testing.T) {
	testUser := fake.NewUser(fake.NewUserEvent("", "name", "email@test.com"))
	id, _ := userService.Create(testUser)
	testUser.ID = *id

	user, _ := userService.FindByID(testUser.ID, nil)
	assert.Equal(t, testUser, *user)
}

func TestUpdate(t *testing.T) {
	testUser := fake.NewUser(fake.NewUserEvent("", "name", "email@test.com"))
	id, _ := userService.Create(testUser)
	testUser.ID = *id
	testUser.Name = "new name"

	doc, _ := util.MapTo[map[string]interface{}](testUser)
	ok, _ := userService.Update(testUser.ID, doc)
	assert.True(t, ok)

	user, _ := userService.FindByID(testUser.ID, nil)
	assert.Equal(t, testUser, *user)
}

func TestDelete(t *testing.T) {
	id, _ := userService.Create(fake.NewUser(fake.NewUserEvent("", "name", "email@test.com")))
	ok, _ := userService.Delete(*id)
	assert.True(t, ok)

	user, _ := userService.FindByID(*id, nil)
	assert.Nil(t, user)
}

func TestTransaction(t *testing.T) {
	id1, _ := accountService.Create(fake.NewAccount(fake.NewAccountEvent("", "name", 50.0, nil)))
	id2, _ := accountService.Create(fake.NewAccount(fake.NewAccountEvent("", "name", 25.0, nil)))
	assert.True(t, !util.NullOrEmpty(id1))
	assert.True(t, !util.NullOrEmpty(id2))

	_, _ = database.Doc("events/event").Set(context.Background(), map[string]interface{}{
		"processed": false,
	})
	assert.False(t, eventService.IsProcessed())

	a1, _ := accountService.FindByID(*id1, nil)
	a2, _ := accountService.FindByID(*id2, nil)
	assert.Equal(t, 50.0, a1.Balance)
	assert.Equal(t, 25.0, a2.Balance)

	createData := map[string]interface{}{"test": "value"}
	err := database.Transaction(func(tx *firestore.Transaction) []data.DatabaseOperation {
		return []data.DatabaseOperation{{
			Ref: accountService.Doc(*id1),
			Data: []firestore.Update{{
				Path:  model.FieldBalance,
				Value: 10.0,
			}},
		}, {
			Ref: accountService.Doc(*id2),
			Data: []firestore.Update{{
				Path:  model.FieldBalance,
				Value: 40.0,
			}},
		}, {
			Ref:    database.Doc("events/transaction-create"),
			Data:   createData,
			Create: true,
		}}
	})

	assert.Nil(t, err)

	a1, _ = accountService.FindByID(*id1, nil)
	a2, _ = accountService.FindByID(*id2, nil)

	snap, _ := database.Doc("events/transaction-create").Get(context.Background())

	assert.Equal(t, 10.0, a1.Balance)
	assert.Equal(t, 40.0, a2.Balance)
	assert.Equal(t, createData, snap.Data())
	assert.True(t, eventService.IsProcessed())
}

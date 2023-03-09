package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/service"
	"finize-functions.app/tests/fake"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestCreate(t *testing.T) {
	database := fake.NewFirestoreService[model.User](context.Background())
	user := fake.NewUser(fake.NewUserEvent("", "name", "email@test.com"))
	id, err := database.Create(service.UsersDB(), nil, fake.MapTo[map[string]interface{}](user))

	assert.True(t, err == nil)
	assert.True(t, !util.NullOrEmpty(id))

	user = fake.NewUser(fake.NewUserEvent("", "name", "email@test.com"))
	id, err = database.Create(service.UsersDB(), util.Pointer("test-id"), fake.MapTo[map[string]interface{}](user))

	assert.True(t, err == nil)
	assert.Equal(t, "test-id", *id)
}

func TestGet(t *testing.T) {
	database := fake.NewFirestoreService[model.User](context.Background())
	testUser := fake.NewUser(fake.NewUserEvent("", "name", "email@test.com"))
	id, _ := database.Create(service.UsersDB(), nil, fake.MapTo[map[string]interface{}](testUser))
	testUser.ID = *id

	user, _ := database.Find(service.UserDoc(testUser.ID), nil)
	assert.True(t, user != nil)
	assert.Equal(t, testUser, *user)
}

func TestGetAll(t *testing.T) {
	database := fake.NewFirestoreService[model.User](context.Background())

	for i := 1; i <= 10; i++ {
		testUser := fake.NewUser(fake.NewUserEvent("", "name", "email@test.com"))
		_, _ = database.Create("get-all-users", util.Pointer("user"+strconv.Itoa(i)), fake.MapTo[map[string]interface{}](testUser))
	}

	users, err := database.GetAll("get-all-users")

	assert.True(t, err == nil)
	assert.Equal(t, 10, len(users))
}

func TestPaginate(t *testing.T) {
	database := fake.NewFirestoreService[model.User](context.Background())

	for i := 10; i <= 30; i++ {
		testUser := fake.NewUser(fake.NewUserEvent("", "name", "email@test.com"))
		_, _ = database.Create("paginate-users", util.Pointer("user"+strconv.Itoa(i)), fake.MapTo[map[string]interface{}](testUser))
	}

	query := database.Collection("paginate-users").OrderBy("id", firestore.Asc)
	page1, err1 := database.Paginate(query, 0, 5)
	page2, err2 := database.Paginate(query, 5, 10)

	assert.True(t, err1 == nil)
	assert.True(t, err2 == nil)
	assert.Equal(t, 5, len(page1))
	assert.Equal(t, 10, len(page2))
	assert.Equal(t, "user10", page1[0].ID)
	assert.Equal(t, "user14", page1[4].ID)
	assert.Equal(t, "user15", page2[0].ID)
	assert.Equal(t, "user24", page2[9].ID)
}

func TestUpdate(t *testing.T) {
	database := fake.NewFirestoreService[model.User](context.Background())
	testUser := fake.NewUser(fake.NewUserEvent("", "name", "email@test.com"))
	id, _ := database.Create(service.UsersDB(), nil, fake.MapTo[map[string]interface{}](testUser))
	testUser.ID = *id
	testUser.Name = "new name"

	ok, _ := database.Update(service.UserDoc(testUser.ID), fake.MapTo[map[string]interface{}](testUser))
	assert.True(t, ok)

	user, _ := database.Find(service.UserDoc(testUser.ID), nil)
	assert.True(t, user != nil)
	assert.Equal(t, testUser, *user)
}

func TestDelete(t *testing.T) {
	database := fake.NewFirestoreService[model.User](context.Background())
	testUser := fake.NewUser(fake.NewUserEvent("", "name", "email@test.com"))
	id, _ := database.Create(service.UsersDB(), nil, fake.MapTo[map[string]interface{}](testUser))
	ok, _ := database.Delete(service.UserDoc(*id))
	assert.True(t, ok)

	user, _ := database.Find(service.UserDoc(*id), nil)
	assert.True(t, user == nil)
}

func TestTransaction(t *testing.T) {
	database := fake.NewFirestoreService[model.Account](context.Background())
	id1, _ := database.Create("accounts", nil, fake.MapTo[map[string]interface{}](fake.NewAccount(fake.NewAccountEvent("", "name", 50.0, nil))))
	id2, _ := database.Create("accounts", nil, fake.MapTo[map[string]interface{}](fake.NewAccount(fake.NewAccountEvent("", "name", 25.0, nil))))
	assert.True(t, !util.NullOrEmpty(id1))
	assert.True(t, !util.NullOrEmpty(id2))

	_, _ = database.Doc(service.EventDoc("event")).Delete(context.Background())
	assert.False(t, eventService.IsProcessed())

	a1, _ := database.Find("accounts/"+*id1, nil)
	a2, _ := database.Find("accounts/"+*id2, nil)
	assert.Equal(t, 50.0, a1.Balance)
	assert.Equal(t, 25.0, a2.Balance)

	createData := map[string]interface{}{"test": "value"}
	err := database.Transaction(func(tx *firestore.Transaction) []data.DatabaseOperation {
		return []data.DatabaseOperation{{
			Ref: database.Doc("accounts/" + *id1),
			Data: []firestore.Update{{
				Path:  model.FieldBalance,
				Value: 10.0,
			}},
		}, {
			Ref: database.Doc("accounts/" + *id2),
			Data: []firestore.Update{{
				Path:  model.FieldBalance,
				Value: 40.0,
			}},
		}, {
			Ref:    database.Doc("accounts/transaction-create"),
			Data:   createData,
			Create: true,
		}}
	})

	assert.True(t, err == nil)

	a1, _ = database.Find("accounts/"+*id1, nil)
	a2, _ = database.Find("accounts/"+*id2, nil)

	snap, _ := database.Doc("accounts/transaction-create").Get(context.Background())

	assert.Equal(t, 10.0, a1.Balance)
	assert.Equal(t, 40.0, a2.Balance)
	assert.Equal(t, createData, snap.Data())
	assert.True(t, eventService.IsProcessed())
}

func TestBatch(t *testing.T) {
	database := fake.NewFirestoreService[model.User](context.Background())

	id1, _ := database.Create(service.UsersDB(), nil, fake.MapTo[map[string]interface{}](fake.NewUser(fake.NewUserEvent("", "name one", "email@test.com"))))
	id2, _ := database.Create(service.UsersDB(), nil, fake.MapTo[map[string]interface{}](fake.NewUser(fake.NewUserEvent("", "name two", "email@test.com"))))
	assert.True(t, !util.NullOrEmpty(id1))
	assert.True(t, !util.NullOrEmpty(id2))

	_, _ = database.Doc(service.EventDoc("event")).Delete(context.Background())
	assert.False(t, eventService.IsProcessed())

	u1, _ := database.Find(service.UserDoc(*id1), nil)
	u2, _ := database.Find(service.UserDoc(*id2), nil)
	assert.Equal(t, "name one", u1.Name)
	assert.Equal(t, "name two", u2.Name)

	createData := map[string]interface{}{"test": "value"}
	err := database.Batch(func() []data.DatabaseOperation {
		return []data.DatabaseOperation{{
			Ref: database.Doc(service.UserDoc(*id1)),
			Data: []firestore.Update{{
				Path:  "name",
				Value: "new name one",
			}},
		}, {
			Ref: database.Doc(service.UserDoc(*id2)),
			Data: []firestore.Update{{
				Path:  "name",
				Value: "new name two",
			}},
		}, {
			Ref:    database.Doc(service.UserDoc("batch-create")),
			Data:   createData,
			Create: true,
		}}
	})

	assert.True(t, err == nil)

	u1, _ = database.Find(service.UserDoc(*id1), nil)
	u2, _ = database.Find(service.UserDoc(*id2), nil)

	snap, _ := database.Doc(service.UserDoc("batch-create")).Get(context.Background())

	assert.Equal(t, "new name one", u1.Name)
	assert.Equal(t, "new name two", u2.Name)
	assert.Equal(t, createData, snap.Data())
	assert.True(t, eventService.IsProcessed())
}

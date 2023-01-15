package data

import (
	"finize-functions.app/data/model"
	"finize-functions.app/tests/fake"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromDocument(t *testing.T) {
	got, _ := util.MapTo[model.UserEvent](fake.NewUserEventMap(fake.NewUserEvent("id", "name", "email")))
	want := fake.NewUserEvent("id", "name", "email")

	assert.Equal(t, want.ID, got.ID)
	assert.Equal(t, want.Name, got.Name)
	assert.Equal(t, want.Email, got.Email)
}

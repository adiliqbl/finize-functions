package tests

import (
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/util"
	"testing"
)

func TestFromDocument(t *testing.T) {
	got, _ := util.FromDocument[model.User](map[string]interface{}{
		"id": map[string]interface{}{
			"stringValue": "id",
		},
		"name": map[string]interface{}{
			"stringValue": "name",
		},
		"email": map[string]interface{}{
			"stringValue": "email",
		},
	})
	want := model.User{
		ID:    data.StringValue{StringValue: "id"},
		Name:  data.StringValue{StringValue: "name"},
		Email: data.StringValue{StringValue: "email"},
	}

	if *got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestToDocument(t *testing.T) {
	got, _ := util.ToDocument(model.User{
		ID:    data.StringValue{StringValue: "id"},
		Name:  data.StringValue{StringValue: "name"},
		Email: data.StringValue{StringValue: "email"},
	})
	want := map[string]interface{}{
		"id":    "id",
		"name":  "name",
		"email": "email",
		"image": "",
	}

	for key, value := range got {
		if want[key] != value {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

	got, _ = util.ToDocument(map[string]interface{}{
		"id":    "id",
		"name":  "name",
		"email": "email",
	})
	want = map[string]interface{}{
		"id":    "id",
		"name":  "name",
		"email": "email",
	}
	for key, value := range got {
		if want[key] != value {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

package finize_functions

import (
	"context"
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/functions"
)

func OnUserCreated(ctx context.Context, e data.FirestoreEvent[model.User]) error {
	return functions.OnUserCreated(ctx, e)
}

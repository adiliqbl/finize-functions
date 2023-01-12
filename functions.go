package finize_functions

import (
	"context"
	"finize-functions/data"
	"finize-functions/data/model"
	"finize-functions/functions"
)

func OnUserCreated(ctx context.Context, e data.FirestoreEvent[model.User]) error {
	return functions.OnUserCreated(ctx, e)
}

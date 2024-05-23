package business

import (
	"context"
	"main/models"
)

func UpdateUserAction(ctx context.Context, ua *models.UserAction) (err error) {
	ua.ID += 1

	return
}

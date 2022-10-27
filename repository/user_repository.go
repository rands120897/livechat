package repository

import (
	"context"
	"webservice/models"
)

type UserRepository interface {
	insert(ctx context.Context, user models.User) (models.User, error)
	update(ctx context.Context, user models.User) (models.User, error)
}

package repository

import (
	"context"

	user "github.com/FelipePn10/panossoerp/internal/domain/user/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *user.User) error
	FindByEmail(ctx context.Context, email string) (*user.User, error)
}

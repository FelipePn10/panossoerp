package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/group/entity"
)

type GroupRepository interface {
	Create(ctx context.Context, group *entity.Group) (*entity.Group, error)
}

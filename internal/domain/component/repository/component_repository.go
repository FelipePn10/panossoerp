package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/component/entity"
	"github.com/google/uuid"
)

type ComponentRepository interface {
	Save(ctx context.Context, component *entity.Component) error
	Delete(ctx context.Context, id int64) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Component, error)
}

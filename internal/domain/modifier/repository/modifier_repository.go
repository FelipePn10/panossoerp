package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/modifier/entity"
)

type ModifierRepository interface {
	Create(ctx context.Context, modifier *entity.Modifier) (*entity.Modifier, error)
}

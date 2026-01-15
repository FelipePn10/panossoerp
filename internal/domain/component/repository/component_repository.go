package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/component/entity"
)

type ComponentRepository interface {
	Save(ctx context.Context, component *entity.Component) (*entity.Component, error)
	ExistsComponentByCode(ctx context.Context, code string) (bool, error)
}

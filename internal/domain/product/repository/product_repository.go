package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/product/entity"
	"github.com/google/uuid"
)

type ProductRepository interface {
	Save(ctx context.Context, product *entity.Product) error
	//FindByID(ctx context.Context, id uuid.UUID) (*entity.Product, error)
}

type ComponentRepository interface {
	Save(ctx context.Context, component *entity.Component) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Component, error)
}

package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/product/entity"
)

type ProductRepository interface {
	Save(ctx context.Context, product *entity.Product) (*entity.Product, error)
	Delete(ctx context.Context, id int64) error
	FindByNameAndCode(ctx context.Context, name string, code string) (*entity.Product, error)
	ExistsByCode(ctx context.Context, code string) (bool, error)
	//FindByID(ctx context.Context, id uuid.UUID) (*entity.Product, error)
}

package product

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/product/entity"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

func (r *repositoryProductSQLC) Save(
	ctx context.Context,
	product *entity.Product,
) error {

	dbProduct, err := r.q.CreateProduct(ctx, sqlc.CreateProductParams{
		ID:        product.ID,
		Code:      product.Code,
		GroupCode: product.GroupCode,
		Name:      product.Name,
		CreatedBy: product.CreatedBy,
	})
	if err != nil {
		return err
	}

	product.CreatedAt = dbProduct.CreatedAt
	product.UpdatedAt = dbProduct.UpdatedAt

	return nil
}

// func (r *repositoryProductSQLC) FindByID(
// 	ctx context.Context,
// 	id uuid.UUID,
// ) (*entity.Product, error) {

// 	dbProduct, err := r.q.GetProductByID(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &entity.Product{
// 		ID:        dbProduct.ID,
// 		Code:      dbProduct.Code,
// 		GroupCode: dbProduct.GroupCode,
// 		Name:      dbProduct.Name,
// 		CreatedBy: dbProduct.CreatedBy,
// 		CreatedAt: dbProduct.CreatedAt,
// 		UpdatedAt: dbProduct.UpdatedAt,
// 	}, nil
// }

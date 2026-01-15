package product

import (
	"context"
	"database/sql"
	"errors"

	"github.com/FelipePn10/panossoerp/internal/domain/product/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/product/repository"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

func (r *repositoryProductSQLC) Save(
	ctx context.Context,
	product *entity.Product,
) (*entity.Product, error) {

	params := sqlc.CreateProductParams{
		ID:   product.ID,
		Code: product.Code,
		GroupCode: sql.NullString{
			String: product.GroupCode,
			Valid:  product.GroupCode != "",
		},
		Name:      product.Name,
		CreatedBy: product.CreatedBy,
	}

	dbProduct, err := r.q.CreateProduct(ctx, params)
	if err != nil {
		return nil, err
	}

	return &entity.Product{
		ID:        dbProduct.ID,
		Code:      dbProduct.Code,
		GroupCode: dbProduct.GroupCode.String,
		Name:      dbProduct.Name,
		CreatedBy: dbProduct.CreatedBy,
		CreatedAt: dbProduct.CreatedAt,
	}, nil

}

func (r *repositoryProductSQLC) Delete(
	ctx context.Context,
	id int64,
) error {
	return r.q.DeleteProduct(ctx, id)
}

func (r *repositoryProductSQLC) ExistsProductByCode(
	ctx context.Context,
	code string,
) (bool, error) {
	_, err := r.q.ExistsProductByCode(ctx, code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *repositoryProductSQLC) FindByNameAndCode(
	ctx context.Context,
	name string,
	code string,
) (*entity.Product, error) {
	dbProduct, err := r.q.FindByNameAndCode(ctx, sqlc.FindByNameAndCodeParams{
		Name: name,
		Code: code,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return &entity.Product{
		ID:        dbProduct.ID,
		Code:      dbProduct.Code,
		GroupCode: dbProduct.GroupCode.String,
		Name:      dbProduct.Name,
		CreatedBy: dbProduct.CreatedBy,
		CreatedAt: dbProduct.CreatedAt,
	}, nil
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

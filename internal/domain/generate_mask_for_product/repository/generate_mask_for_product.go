package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/generate_mask_for_product/entity"
)

type GenerateMaskForProductRepository interface {
	Generate(ctx context.Context, mask *entity.ProductMask) error
}

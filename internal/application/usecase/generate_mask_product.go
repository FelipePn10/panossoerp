package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/generate_mask_for_product/repository"
)

type GenerateMaskForProduct struct {
	repo repository.GenerateMaskForProductRepository
}

func (uc *GenerateMaskForProduct) Execute(
	ctx context.Context,
	dto repository.GenerateMaskForProductRepository,
) error {

}

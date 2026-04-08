package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/generate_mask_for_item/entity"
)

type GenerateMaskForItemRepository interface {
	Generate(ctx context.Context, mask *entity.ItemMask) (*entity.ItemMask, error)
	GetOptionValue(ctx context.Context, optionID int64) (string, error)
}

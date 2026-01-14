package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/bom/entity"
)

type BomRepository interface {
	Create(ctx context.Context, bom *entity.Bom) (*entity.Bom, error)
	ExistsByID(ctx context.Context, id int64) (bool, error)
}

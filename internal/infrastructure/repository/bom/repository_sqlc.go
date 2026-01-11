package bom

import (
	"context"
	"database/sql"

	"github.com/FelipePn10/panossoerp/internal/domain/bom/entity"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

func (r *repositoryBomSQLC) Create(
	ctx context.Context,
	bom *entity.Bom,
) (*entity.Bom, error) {
	params := sqlc.CreateBomParams{
		ProductID: bom.ProductId,
		BomType:   bom.BomType,
		Version:   int32(bom.Version),
		ValidFrom: sql.NullTime{},
		Status:    bom.Status,
	}
	_, err := r.q.CreateBom(ctx, params)
	if err != nil {
		return nil, err
	}

	return bom, nil
}

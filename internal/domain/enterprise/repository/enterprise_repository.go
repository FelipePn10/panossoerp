package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/enterprise/entity"
)

type EnterpriseRepository interface {
	Create(ctx context.Context, enterprise *entity.Enterprise) (*entity.Enterprise, error)
}

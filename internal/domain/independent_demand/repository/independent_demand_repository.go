package repository

import (
	"context"
	"time"

	"github.com/FelipePn10/panossoerp/internal/domain/independent_demand/entity"
)

type IndependentDemandRepository interface {
	Create(ctx context.Context, d *entity.IndependentDemand) (*entity.IndependentDemand, error)
	Update(ctx context.Context, d *entity.IndependentDemand) (*entity.IndependentDemand, error)
	GetByCode(ctx context.Context, code int64) (*entity.IndependentDemand, error)
	List(ctx context.Context) ([]*entity.IndependentDemand, error)
	ListByItem(ctx context.Context, itemCode int64) ([]*entity.IndependentDemand, error)
	ListFromDate(ctx context.Context, date time.Time) ([]*entity.IndependentDemand, error)
	Delete(ctx context.Context, code int64) error
}

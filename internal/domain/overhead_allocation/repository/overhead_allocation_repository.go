package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/overhead_allocation/entity"
)

type OverheadAllocationRepository interface {
	Create(ctx context.Context, oa *entity.OverheadAllocation) (*entity.OverheadAllocation, error)
	AddTarget(ctx context.Context, target *entity.AllocationTarget) (*entity.AllocationTarget, error)
	GetByCode(ctx context.Context, code int64) (*entity.OverheadAllocation, error)
	GetTargets(ctx context.Context, overheadCode int64) ([]*entity.AllocationTarget, error)
	List(ctx context.Context) ([]*entity.OverheadAllocation, error)
	ListByCostCenter(ctx context.Context, ccCode int64) ([]*entity.OverheadAllocation, error)
	Delete(ctx context.Context, code int64) error
	DeleteTargets(ctx context.Context, overheadCoed int64) error
}

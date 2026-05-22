package repository

import (
	"context"
	"time"

	"github.com/FelipePn10/panossoerp/internal/domain/maintenance/entity"
)

type MaintenanceRepository interface {
	// Plans
	CreatePlan(ctx context.Context, p *entity.MaintenancePlan) (*entity.MaintenancePlan, error)
	UpdatePlan(ctx context.Context, p *entity.MaintenancePlan) (*entity.MaintenancePlan, error)
	GetPlanByID(ctx context.Context, id int64) (*entity.MaintenancePlan, error)
	ListPlans(ctx context.Context, onlyActive bool) ([]*entity.MaintenancePlan, error)
	ListPlansByMachine(ctx context.Context, machineID int64) ([]*entity.MaintenancePlan, error)
	DeactivatePlan(ctx context.Context, id int64) error

	// Orders
	CreateOrder(ctx context.Context, o *entity.MaintenanceOrder) (*entity.MaintenanceOrder, error)
	UpdateOrder(ctx context.Context, o *entity.MaintenanceOrder) (*entity.MaintenanceOrder, error)
	GetOrderByID(ctx context.Context, id int64) (*entity.MaintenanceOrder, error)
	ListOrdersByPlan(ctx context.Context, planID int64) ([]*entity.MaintenanceOrder, error)
	ListOrdersByWorkCenter(ctx context.Context, workCenterID int64, from, to time.Time) ([]*entity.MaintenanceOrder, error)

	// ExistsOrderForPlanAndDate returns true when a non-cancelled order already
	// exists for this plan on the given date — used to prevent duplicates in GenerateOrders.
	ExistsOrderForPlanAndDate(ctx context.Context, planID int64, date time.Time) (bool, error)

	// CRP integration: returns total maintenance hours scheduled for a work center on a date range
	GetBlockedHours(ctx context.Context, workCenterID int64, date time.Time) (float64, error)
}

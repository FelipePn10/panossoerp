package repository

import (
	"context"
	"time"

	"github.com/FelipePn10/panossoerp/internal/domain/aps/entity"
)

type APSRepository interface {
	UpsertSequence(ctx context.Context, seq *entity.ProductionSequence) (*entity.ProductionSequence, error)
	ListByOrder(ctx context.Context, orderID int64) ([]*entity.ProductionSequence, error)
	ListByWorkCenter(ctx context.Context, workCenterID int64, from, to time.Time) ([]*entity.ProductionSequence, error)
	DeleteByOrder(ctx context.Context, orderID int64) error

	// Data needed by the sequencing algorithm
	GetOpenProductionOrders(ctx context.Context) ([]OrderRow, error)
	GetOrderOperations(ctx context.Context, orderID int64) ([]OpRow, error)
	GetWorkCenterCapacity(ctx context.Context, workCenterID int64) (float64, error)
}

type OrderRow struct {
	ID          int64
	Priority    int
	PlannedDate time.Time
}

type OpRow struct {
	ID           int64
	Sequence     int
	WorkCenterID *int64
	PlannedHours float64
	SetupHours   float64
}

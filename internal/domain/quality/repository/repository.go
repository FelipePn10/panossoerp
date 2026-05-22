package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/quality/entity"
)

type QualityRepository interface {
	// Inspection plans
	CreatePlan(ctx context.Context, plan *entity.InspectionPlan) (*entity.InspectionPlan, error)
	GetPlanByID(ctx context.Context, id int64) (*entity.InspectionPlan, error)
	ListPlansByItem(ctx context.Context, itemCode int64) ([]*entity.InspectionPlan, error)
	DeactivatePlan(ctx context.Context, id int64) error

	// Characteristics
	AddCharacteristic(ctx context.Context, c *entity.InspectionCharacteristic) (*entity.InspectionCharacteristic, error)
	ListCharacteristics(ctx context.Context, planID int64) ([]*entity.InspectionCharacteristic, error)

	// Quality records
	CreateRecord(ctx context.Context, r *entity.QualityRecord) (*entity.QualityRecord, error)
	AddMeasurement(ctx context.Context, m *entity.QualityMeasurement) error
	GetRecordByID(ctx context.Context, id int64) (*entity.QualityRecord, error)
	ListRecordsByOrder(ctx context.Context, orderID int64) ([]*entity.QualityRecord, error)
	ListRecordsByItem(ctx context.Context, itemCode int64) ([]*entity.QualityRecord, error)

	// Non-conformances
	CreateNC(ctx context.Context, nc *entity.NonConformance) (*entity.NonConformance, error)
	GetNCByID(ctx context.Context, id int64) (*entity.NonConformance, error)
	ListOpenNCs(ctx context.Context) ([]*entity.NonConformance, error)
	ListNCsByItem(ctx context.Context, itemCode int64) ([]*entity.NonConformance, error)
	DispositionNC(ctx context.Context, id int64, disposition entity.NCDisposition, disposedBy string) error
	NextNCCode(ctx context.Context) (int64, error)
}

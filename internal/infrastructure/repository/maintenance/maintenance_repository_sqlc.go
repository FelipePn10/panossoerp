package maintenance

import (
	"context"
	"fmt"
	"time"

	"github.com/FelipePn10/panossoerp/internal/domain/maintenance/entity"
	domainrepo "github.com/FelipePn10/panossoerp/internal/domain/maintenance/repository"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/pgutil"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type MaintenanceRepositorySQLC struct {
	q *sqlc.Queries
}

func New(q *sqlc.Queries) domainrepo.MaintenanceRepository {
	return &MaintenanceRepositorySQLC{q: q}
}

// ─── plans ────────────────────────────────────────────────────────────────────

func (r *MaintenanceRepositorySQLC) CreatePlan(ctx context.Context, p *entity.MaintenancePlan) (*entity.MaintenancePlan, error) {
	nextSched := pgutil.ToPgTimestamptz(time.Now().AddDate(0, 0, p.FrequencyDays))
	if p.NextScheduledAt != nil {
		nextSched = pgutil.ToPgTimestamptz(*p.NextScheduledAt)
	}
	row, err := r.q.CreateMaintenancePlan(ctx, sqlc.CreateMaintenancePlanParams{
		MachineID:       p.MachineID,
		WorkCenterID:    pgutil.ToPgInt8Ptr(p.WorkCenterID),
		Description:     p.Description,
		Frequency:       sqlc.MaintenanceFrequencyEnum(p.Frequency),
		FrequencyDays:   int32(p.FrequencyDays),
		EstimatedHours:  p.EstimatedHours,
		NextScheduledAt: nextSched,
		CreatedBy:       pgutil.ToPgUUID(p.CreatedBy),
	})
	if err != nil {
		return nil, fmt.Errorf("creating maintenance plan: %w", err)
	}
	return planRowToEntity(row), nil
}

func (r *MaintenanceRepositorySQLC) UpdatePlan(ctx context.Context, p *entity.MaintenancePlan) (*entity.MaintenancePlan, error) {
	nextSched := pgutil.ToPgTimestamptz(time.Now().AddDate(0, 0, p.FrequencyDays))
	if p.NextScheduledAt != nil {
		nextSched = pgutil.ToPgTimestamptz(*p.NextScheduledAt)
	}
	row, err := r.q.UpdateMaintenancePlan(ctx, sqlc.UpdateMaintenancePlanParams{
		ID:              p.ID,
		Description:     p.Description,
		Frequency:       sqlc.MaintenanceFrequencyEnum(p.Frequency),
		FrequencyDays:   int32(p.FrequencyDays),
		EstimatedHours:  p.EstimatedHours,
		NextScheduledAt: nextSched,
	})
	if err != nil {
		return nil, fmt.Errorf("updating maintenance plan: %w", err)
	}
	return planRowToEntity(row), nil
}

func (r *MaintenanceRepositorySQLC) GetPlanByID(ctx context.Context, id int64) (*entity.MaintenancePlan, error) {
	row, err := r.q.GetMaintenancePlanByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("fetching maintenance plan %d: %w", id, err)
	}
	return planRowToEntity(row), nil
}

func (r *MaintenanceRepositorySQLC) ListPlans(ctx context.Context, onlyActive bool) ([]*entity.MaintenancePlan, error) {
	rows, err := r.q.ListMaintenancePlans(ctx, &onlyActive)
	if err != nil {
		return nil, fmt.Errorf("listing maintenance plans: %w", err)
	}
	return planSlice(rows), nil
}

func (r *MaintenanceRepositorySQLC) ListPlansByMachine(ctx context.Context, machineID int64) ([]*entity.MaintenancePlan, error) {
	rows, err := r.q.ListMaintenancePlansByMachine(ctx, machineID)
	if err != nil {
		return nil, fmt.Errorf("listing plans for machine %d: %w", machineID, err)
	}
	return planSlice(rows), nil
}

func (r *MaintenanceRepositorySQLC) DeactivatePlan(ctx context.Context, id int64) error {
	return r.q.DeactivateMaintenancePlan(ctx, id)
}

// ─── orders ───────────────────────────────────────────────────────────────────

func (r *MaintenanceRepositorySQLC) CreateOrder(ctx context.Context, o *entity.MaintenanceOrder) (*entity.MaintenanceOrder, error) {
	row, err := r.q.CreateMaintenanceOrder(ctx, sqlc.CreateMaintenanceOrderParams{
		PlanID:         o.PlanID,
		MachineID:      pgutil.ToPgInt8Ptr(o.MachineID),
		WorkCenterID:   pgutil.ToPgInt8Ptr(o.WorkCenterID),
		ScheduledDate:  pgutil.ToPgDate(o.ScheduledDate),
		EstimatedHours: o.EstimatedHours,
	})
	if err != nil {
		return nil, fmt.Errorf("creating maintenance order: %w", err)
	}
	return orderRowToEntity(row), nil
}

func (r *MaintenanceRepositorySQLC) UpdateOrder(ctx context.Context, o *entity.MaintenanceOrder) (*entity.MaintenanceOrder, error) {
	actualH := pgtype.Float8{}
	if o.ActualHours != nil {
		actualH = pgtype.Float8{Float64: *o.ActualHours, Valid: true}
	}
	startedAt := pgtype.Timestamptz{}
	if o.StartedAt != nil {
		startedAt = pgutil.ToPgTimestamptz(*o.StartedAt)
	}
	completedAt := pgtype.Timestamptz{}
	if o.CompletedAt != nil {
		completedAt = pgutil.ToPgTimestamptz(*o.CompletedAt)
	}
	row, err := r.q.UpdateMaintenanceOrder(ctx, sqlc.UpdateMaintenanceOrderParams{
		ID:          o.ID,
		Status:      sqlc.MaintenanceOrderStatusEnum(o.Status),
		ActualHours: actualH,
		StartedAt:   startedAt,
		CompletedAt: completedAt,
		Notes:       pgutil.ToPgTextFromPtr(o.Notes),
	})
	if err != nil {
		return nil, fmt.Errorf("updating maintenance order %d: %w", o.ID, err)
	}
	return orderRowToEntity(row), nil
}

func (r *MaintenanceRepositorySQLC) GetOrderByID(ctx context.Context, id int64) (*entity.MaintenanceOrder, error) {
	row, err := r.q.GetMaintenanceOrderByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("fetching maintenance order %d: %w", id, err)
	}
	return orderRowToEntity(row), nil
}

func (r *MaintenanceRepositorySQLC) ListOrdersByPlan(ctx context.Context, planID int64) ([]*entity.MaintenanceOrder, error) {
	rows, err := r.q.ListMaintenanceOrdersByPlan(ctx, planID)
	if err != nil {
		return nil, fmt.Errorf("listing orders for plan %d: %w", planID, err)
	}
	return orderSlice(rows), nil
}

func (r *MaintenanceRepositorySQLC) ListOrdersByWorkCenter(ctx context.Context, workCenterID int64, from, to time.Time) ([]*entity.MaintenanceOrder, error) {
	rows, err := r.q.ListMaintenanceOrdersByWorkCenter(ctx, workCenterID,
		pgutil.ToPgDate(from), pgutil.ToPgDate(to))
	if err != nil {
		return nil, fmt.Errorf("listing orders for work center %d: %w", workCenterID, err)
	}
	return orderSlice(rows), nil
}

func (r *MaintenanceRepositorySQLC) ExistsOrderForPlanAndDate(ctx context.Context, planID int64, date time.Time) (bool, error) {
	return r.q.ExistsOrderForPlanAndDate(ctx, planID, date)
}

func (r *MaintenanceRepositorySQLC) GetBlockedHours(ctx context.Context, workCenterID int64, date time.Time) (float64, error) {
	return r.q.GetBlockedHoursOnDate(ctx, workCenterID, date)
}

// ─── mappers ──────────────────────────────────────────────────────────────────

func planRowToEntity(row sqlc.DBMaintenancePlan) *entity.MaintenancePlan {
	p := &entity.MaintenancePlan{
		ID:             row.ID,
		Code:           row.Code,
		MachineID:      row.MachineID,
		Description:    row.Description,
		Frequency:      entity.Frequency(row.Frequency),
		FrequencyDays:  int(row.FrequencyDays),
		EstimatedHours: row.EstimatedHours,
		IsActive:       row.IsActive,
		CreatedAt:      pgutil.FromPgTimestamptz(row.CreatedAt),
		UpdatedAt:      pgutil.FromPgTimestamptz(row.UpdatedAt),
		CreatedBy:      pgutil.FromPgUUID(row.CreatedBy),
	}
	if row.WorkCenterID.Valid {
		v := row.WorkCenterID.Int64
		p.WorkCenterID = &v
	}
	if row.LastExecutedAt.Valid {
		t := pgutil.FromPgTimestamptz(row.LastExecutedAt)
		p.LastExecutedAt = &t
	}
	if row.NextScheduledAt.Valid {
		t := pgutil.FromPgTimestamptz(row.NextScheduledAt)
		p.NextScheduledAt = &t
	}
	return p
}

func planSlice(rows []sqlc.DBMaintenancePlan) []*entity.MaintenancePlan {
	out := make([]*entity.MaintenancePlan, 0, len(rows))
	for _, row := range rows {
		out = append(out, planRowToEntity(row))
	}
	return out
}

func orderRowToEntity(row sqlc.DBMaintenanceOrder) *entity.MaintenanceOrder {
	o := &entity.MaintenanceOrder{
		ID:             row.ID,
		PlanID:         row.PlanID,
		ScheduledDate:  pgutil.FromPgDate(row.ScheduledDate),
		EstimatedHours: row.EstimatedHours,
		Status:         entity.OrderStatus(row.Status),
		IsActive:       row.IsActive,
		CreatedAt:      pgutil.FromPgTimestamptz(row.CreatedAt),
		UpdatedAt:      pgutil.FromPgTimestamptz(row.UpdatedAt),
	}
	if row.MachineID.Valid {
		v := row.MachineID.Int64
		o.MachineID = &v
	}
	if row.WorkCenterID.Valid {
		v := row.WorkCenterID.Int64
		o.WorkCenterID = &v
	}
	if row.ActualHours.Valid {
		v := row.ActualHours.Float64
		o.ActualHours = &v
	}
	if row.StartedAt.Valid {
		t := pgutil.FromPgTimestamptz(row.StartedAt)
		o.StartedAt = &t
	}
	if row.CompletedAt.Valid {
		t := pgutil.FromPgTimestamptz(row.CompletedAt)
		o.CompletedAt = &t
	}
	if row.Notes.Valid {
		v := row.Notes.String
		o.Notes = &v
	}
	return o
}

func orderSlice(rows []sqlc.DBMaintenanceOrder) []*entity.MaintenanceOrder {
	out := make([]*entity.MaintenanceOrder, 0, len(rows))
	for _, row := range rows {
		out = append(out, orderRowToEntity(row))
	}
	return out
}

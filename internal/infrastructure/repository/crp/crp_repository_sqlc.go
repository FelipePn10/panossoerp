package crp

import (
	"context"
	"fmt"
	"time"

	"github.com/FelipePn10/panossoerp/internal/domain/crp/entity"
	domainrepo "github.com/FelipePn10/panossoerp/internal/domain/crp/repository"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/pgutil"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

type CRPRepositorySQLC struct {
	q *sqlc.Queries
}

func New(q *sqlc.Queries) domainrepo.CRPRepository {
	return &CRPRepositorySQLC{q: q}
}

func (r *CRPRepositorySQLC) UpsertRequirement(ctx context.Context, req *entity.CapacityRequirement) (*entity.CapacityRequirement, error) {
	row, err := r.q.UpsertCapacityRequirement(ctx, sqlc.UpsertCapacityRequirementParams{
		PlanCode:       req.PlanCode,
		WorkCenterID:   req.WorkCenterID,
		ReqDate:        sqlc.ToPgDate(req.ReqDate),
		RequiredHours:  pgutil.ToPgNumericFromFloat64(req.RequiredHours),
		AvailableHours: pgutil.ToPgNumericFromFloat64(req.AvailableHours),
	})
	if err != nil {
		return nil, fmt.Errorf("upserting capacity requirement: %w", err)
	}
	return crpRowToEntity(row), nil
}

func (r *CRPRepositorySQLC) ListByPlan(ctx context.Context, planCode int64) ([]*entity.CapacityRequirement, error) {
	rows, err := r.q.ListCRPByPlan(ctx, planCode)
	if err != nil {
		return nil, fmt.Errorf("listing CRP for plan %d: %w", planCode, err)
	}
	return crpSlice(rows), nil
}

func (r *CRPRepositorySQLC) ListOverloadedByPlan(ctx context.Context, planCode int64) ([]*entity.CapacityRequirement, error) {
	rows, err := r.q.ListOverloadedCRPByPlan(ctx, planCode)
	if err != nil {
		return nil, fmt.Errorf("listing overloaded CRP for plan %d: %w", planCode, err)
	}
	return crpSlice(rows), nil
}

func (r *CRPRepositorySQLC) ListByWorkCenter(ctx context.Context, workCenterID int64, from, to time.Time) ([]*entity.CapacityRequirement, error) {
	rows, err := r.q.ListCRPByWorkCenter(ctx, workCenterID, sqlc.ToPgDate(from), sqlc.ToPgDate(to))
	if err != nil {
		return nil, fmt.Errorf("listing CRP for work center %d: %w", workCenterID, err)
	}
	return crpSlice(rows), nil
}

func (r *CRPRepositorySQLC) DeleteByPlan(ctx context.Context, planCode int64) error {
	return r.q.DeleteCRPByPlan(ctx, planCode)
}

func (r *CRPRepositorySQLC) GetPlannedOrdersByPlan(ctx context.Context, planCode int64) ([]domainrepo.PlannedOrderRow, error) {
	rows, err := r.q.GetPlannedOrdersForCRP(ctx, planCode)
	if err != nil {
		return nil, fmt.Errorf("fetching planned orders for CRP: %w", err)
	}
	out := make([]domainrepo.PlannedOrderRow, 0, len(rows))
	for _, row := range rows {
		po := domainrepo.PlannedOrderRow{
			ID:          row.ID,
			ItemCode:    row.ItemCode,
			Quantity:    row.Quantity,
			PlannedDate: sqlc.FromPgDate(row.PlannedDate),
		}
		if row.RouteID.Valid {
			v := row.RouteID.Int64
			po.RouteID = &v
		}
		out = append(out, po)
	}
	return out, nil
}

func (r *CRPRepositorySQLC) GetRouteOperationsByRoute(ctx context.Context, routeID int64) ([]domainrepo.RouteOpRow, error) {
	rows, err := r.q.GetRouteOpHoursForCRP(ctx, routeID)
	if err != nil {
		return nil, fmt.Errorf("fetching route operations for CRP: %w", err)
	}
	out := make([]domainrepo.RouteOpRow, 0, len(rows))
	for _, row := range rows {
		ro := domainrepo.RouteOpRow{EffHours: row.EffHours}
		if row.WorkCenterID.Valid {
			v := row.WorkCenterID.Int64
			ro.WorkCenterID = &v
		}
		out = append(out, ro)
	}
	return out, nil
}

func (r *CRPRepositorySQLC) GetMachineAvailableHoursPerDay(ctx context.Context, workCenterID int64) (float64, error) {
	return r.q.GetMachineAvailableHours(ctx, workCenterID)
}

// ─── mappers ──────────────────────────────────────────────────────────────────

func crpRowToEntity(row sqlc.DBCapacityRequirement) *entity.CapacityRequirement {
	reqHours := pgutil.FromPgNumericToFloat64(row.RequiredHours)
	availHours := pgutil.FromPgNumericToFloat64(row.AvailableHours)
	loadPct := pgutil.FromPgNumericToFloat64(row.LoadPct)
	return &entity.CapacityRequirement{
		ID:             row.ID,
		PlanCode:       row.PlanCode,
		WorkCenterID:   row.WorkCenterID,
		ReqDate:        sqlc.FromPgDate(row.ReqDate),
		RequiredHours:  reqHours,
		AvailableHours: availHours,
		LoadPct:        loadPct,
		CreatedAt:      pgutil.FromPgTimestamptz(row.CreatedAt),
	}
}

func crpSlice(rows []sqlc.DBCapacityRequirement) []*entity.CapacityRequirement {
	out := make([]*entity.CapacityRequirement, 0, len(rows))
	for _, row := range rows {
		out = append(out, crpRowToEntity(row))
	}
	return out
}

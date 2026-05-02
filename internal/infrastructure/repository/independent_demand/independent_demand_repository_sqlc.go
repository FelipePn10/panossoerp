package independent_demand

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/FelipePn10/panossoerp/internal/domain/independent_demand/entity"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

func (r *IndependentDemandRepositorySQLC) Create(ctx context.Context,
	d *entity.IndependentDemand) (*entity.IndependentDemand, error) {
	row, err := r.q.CreateIndependentDemand(
		ctx, sqlc.CreateIndependentDemandParams{
			Code:           d.CodeDemand,
			ItemCode:       d.ItemCode,
			Mask:           toNullString(d.Mask),
			CostCenterCode: toNullInt64(d.CostCenterCode),
			Quantity:       strconv.FormatFloat(d.Quantity, 'f', 4, 64),
			DemandDate:     d.DemandDate,
			CreatedBy:      d.CreatedBy,
		})
	if err != nil {
		return nil, fmt.Errorf("creating independent demand: %w", err)
	}
	return rowToEntity(row), nil
}

func (r *IndependentDemandRepositorySQLC) Update(ctx context.Context, d *entity.IndependentDemand) (*entity.IndependentDemand, error) {
	row, err := r.q.UpdateIndependentDemand(ctx, sqlc.UpdateIndependentDemandParams{
		ItemCode:       d.ItemCode,
		Mask:           toNullString(d.Mask),
		CostCenterCode: toNullInt64(d.CostCenterCode),
		Quantity:       strconv.FormatFloat(d.Quantity, 'f', 4, 64),
		DemandDate:     d.DemandDate,
		Code:           d.CodeDemand,
	})
	if err != nil {
		return nil, fmt.Errorf("updating independent demand: %w", err)
	}
	return rowToEntity(row), nil
}

func (r *IndependentDemandRepositorySQLC) GetByCode(ctx context.Context, code int64) (*entity.IndependentDemand, error) {
	row, err := r.q.GetIndependentDemandByCode(ctx, code)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("independent demand code %d not found", code)
		}
		return nil, fmt.Errorf("fetching independent demand by code: %w", err)
	}
	return rowToEntity(row), nil
}

func (r *IndependentDemandRepositorySQLC) List(ctx context.Context) ([]*entity.IndependentDemand, error) {
	rows, err := r.q.ListIndependentDemands(ctx)
	if err != nil {
		return nil, fmt.Errorf("listing independent demands: %w", err)
	}
	return rowsToEntities(rows), nil
}

func (r *IndependentDemandRepositorySQLC) ListByItem(ctx context.Context, itemCode int64) ([]*entity.IndependentDemand, error) {
	rows, err := r.q.ListDemandsByItem(ctx, itemCode)
	if err != nil {
		return nil, fmt.Errorf("listing demands by item: %w", err)
	}
	return rowsToEntities(rows), nil
}

func (r *IndependentDemandRepositorySQLC) ListFromDate(ctx context.Context, date time.Time) ([]*entity.IndependentDemand, error) {
	rows, err := r.q.ListDemandsFromDate(ctx, date)
	if err != nil {
		return nil, fmt.Errorf("listing demands from date: %w", err)
	}
	return rowsToEntities(rows), nil
}

func (r *IndependentDemandRepositorySQLC) Delete(ctx context.Context, id int64) error {
	return r.q.DeleteIndependentDemand(ctx, id)
}

func rowToEntity(row sqlc.IndependentDemand) *entity.IndependentDemand {
	qty, err := strconv.ParseFloat(row.Quantity, 64)
	if err != nil {
		panic(fmt.Sprintf("invalid quantity from database: %v", err))
	}
	e := &entity.IndependentDemand{
		CodeDemand: row.Code,
		ItemCode:   row.ItemCode,
		Quantity:   qty,
		DemandDate: row.DemandDate,
		IsActive:   row.IsActive,
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
		CreatedBy:  row.CreatedBy,
	}
	if row.Mask.Valid {
		v := row.Mask.String
		e.Mask = &v
	}
	if row.CostCenterCode.Valid {
		v := row.CostCenterCode.Int64
		e.CostCenterCode = &v
	}
	return e
}

func rowsToEntities(rows []sqlc.IndependentDemand) []*entity.IndependentDemand {
	out := make([]*entity.IndependentDemand, 0, len(rows))
	for _, row := range rows {
		out = append(out, rowToEntity(row))
	}
	return out
}

func toNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *s, Valid: true}
}

func toNullInt64(i *int64) sql.NullInt64 {
	if i == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{Int64: *i, Valid: true}
}

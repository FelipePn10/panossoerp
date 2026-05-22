package maintenance_uc

import (
	"context"
	"fmt"
	"time"

	"github.com/FelipePn10/panossoerp/internal/domain/maintenance/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/maintenance/repository"
	"github.com/google/uuid"
)

type MaintenanceUseCase struct {
	repo repository.MaintenanceRepository
}

func New(repo repository.MaintenanceRepository) *MaintenanceUseCase {
	return &MaintenanceUseCase{repo: repo}
}

// ─── plans ────────────────────────────────────────────────────────────────────

type CreatePlanDTO struct {
	MachineID      int64   `json:"machine_id"`
	WorkCenterID   *int64  `json:"work_center_id,omitempty"`
	Description    string  `json:"description"`
	Frequency      string  `json:"frequency"`
	FrequencyDays  int     `json:"frequency_days"`
	EstimatedHours float64 `json:"estimated_hours"`
	CreatedBy      string  `json:"created_by"`
}

func (uc *MaintenanceUseCase) CreatePlan(ctx context.Context, dto CreatePlanDTO) (*entity.MaintenancePlan, error) {
	createdBy, err := uuid.Parse(dto.CreatedBy)
	if err != nil {
		return nil, fmt.Errorf("invalid created_by UUID: %w", err)
	}
	plan, err := entity.NewMaintenancePlan(
		dto.MachineID,
		dto.WorkCenterID,
		dto.Description,
		entity.Frequency(dto.Frequency),
		dto.FrequencyDays,
		dto.EstimatedHours,
		createdBy,
	)
	if err != nil {
		return nil, err
	}
	return uc.repo.CreatePlan(ctx, plan)
}

func (uc *MaintenanceUseCase) GetPlan(ctx context.Context, id int64) (*entity.MaintenancePlan, error) {
	return uc.repo.GetPlanByID(ctx, id)
}

func (uc *MaintenanceUseCase) ListPlans(ctx context.Context, onlyActive bool) ([]*entity.MaintenancePlan, error) {
	return uc.repo.ListPlans(ctx, onlyActive)
}

func (uc *MaintenanceUseCase) ListPlansByMachine(ctx context.Context, machineID int64) ([]*entity.MaintenancePlan, error) {
	return uc.repo.ListPlansByMachine(ctx, machineID)
}

func (uc *MaintenanceUseCase) DeactivatePlan(ctx context.Context, id int64) error {
	return uc.repo.DeactivatePlan(ctx, id)
}

// ─── orders ───────────────────────────────────────────────────────────────────

type CreateOrderDTO struct {
	PlanID         int64   `json:"plan_id"`
	WorkCenterID   *int64  `json:"work_center_id,omitempty"`
	ScheduledDate  string  `json:"scheduled_date"` // "2006-01-02"
	EstimatedHours float64 `json:"estimated_hours"`
}

func (uc *MaintenanceUseCase) CreateOrder(ctx context.Context, dto CreateOrderDTO) (*entity.MaintenanceOrder, error) {
	scheduledDate, err := time.Parse("2006-01-02", dto.ScheduledDate)
	if err != nil {
		return nil, fmt.Errorf("invalid scheduled_date: %w", err)
	}
	plan, err := uc.repo.GetPlanByID(ctx, dto.PlanID)
	if err != nil {
		return nil, fmt.Errorf("plan %d not found: %w", dto.PlanID, err)
	}
	workCenterID := dto.WorkCenterID
	if workCenterID == nil {
		workCenterID = plan.WorkCenterID
	}
	estimatedHours := dto.EstimatedHours
	if estimatedHours <= 0 {
		estimatedHours = plan.EstimatedHours
	}
	machineID := plan.MachineID
	order, err := entity.NewMaintenanceOrder(dto.PlanID, &machineID, workCenterID, scheduledDate, estimatedHours)
	if err != nil {
		return nil, err
	}
	return uc.repo.CreateOrder(ctx, order)
}

type AdvanceOrderDTO struct {
	OrderID     int64    `json:"order_id"`
	Status      string   `json:"status"`
	ActualHours *float64 `json:"actual_hours,omitempty"`
	Notes       *string  `json:"notes,omitempty"`
}

func (uc *MaintenanceUseCase) AdvanceOrder(ctx context.Context, dto AdvanceOrderDTO) (*entity.MaintenanceOrder, error) {
	order, err := uc.repo.GetOrderByID(ctx, dto.OrderID)
	if err != nil {
		return nil, fmt.Errorf("order %d not found: %w", dto.OrderID, err)
	}

	order.Status = entity.OrderStatus(dto.Status)
	order.ActualHours = dto.ActualHours
	order.Notes = dto.Notes

	now := time.Now()
	switch order.Status {
	case entity.OrderStatusInProgress:
		if order.StartedAt == nil {
			order.StartedAt = &now
		}
	case entity.OrderStatusDone:
		order.CompletedAt = &now
	}

	return uc.repo.UpdateOrder(ctx, order)
}

func (uc *MaintenanceUseCase) ListOrdersByPlan(ctx context.Context, planID int64) ([]*entity.MaintenanceOrder, error) {
	return uc.repo.ListOrdersByPlan(ctx, planID)
}

func (uc *MaintenanceUseCase) ListOrdersByWorkCenter(ctx context.Context, workCenterID int64, from, to time.Time) ([]*entity.MaintenanceOrder, error) {
	return uc.repo.ListOrdersByWorkCenter(ctx, workCenterID, from, to)
}

// GenerateOrders auto-creates maintenance orders for all active plans whose
// next_scheduled_at falls within the planning horizon (today + horizonDays).
func (uc *MaintenanceUseCase) GenerateOrders(ctx context.Context, horizonDays int) (int, error) {
	plans, err := uc.repo.ListPlans(ctx, true)
	if err != nil {
		return 0, err
	}
	horizon := time.Now().AddDate(0, 0, horizonDays)
	created := 0
	for _, plan := range plans {
		if plan.NextScheduledAt == nil || plan.NextScheduledAt.After(horizon) {
			continue
		}
		scheduled := *plan.NextScheduledAt
		exists, err := uc.repo.ExistsOrderForPlanAndDate(ctx, plan.ID, scheduled)
		if err != nil || exists {
			continue
		}
		order, err := entity.NewMaintenanceOrder(
			plan.ID, &plan.MachineID, plan.WorkCenterID, scheduled, plan.EstimatedHours,
		)
		if err != nil {
			continue
		}
		if _, err := uc.repo.CreateOrder(ctx, order); err == nil {
			created++
		}
	}
	return created, nil
}

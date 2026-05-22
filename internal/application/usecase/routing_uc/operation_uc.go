package routing_uc

import (
	"context"
	"fmt"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/application/dto/response"
	"github.com/FelipePn10/panossoerp/internal/domain/routing/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/routing/repository"
)

type OperationUseCase struct {
	repo repository.RoutingRepository
}

func NewOperationUseCase(repo repository.RoutingRepository) *OperationUseCase {
	return &OperationUseCase{repo: repo}
}

func (uc *OperationUseCase) Create(ctx context.Context, dto request.CreateOperationDTO) (*response.OperationResponse, error) {
	if dto.Name == "" {
		return nil, fmt.Errorf("name is required")
	}
	origin := entity.OperationOrigin(dto.Origin)
	if origin == "" {
		origin = entity.OriginInternal
	}

	code, err := uc.repo.NextOperationCode(ctx)
	if err != nil {
		return nil, fmt.Errorf("generating operation code: %w", err)
	}

	op, err := entity.NewOperation(code, dto.Name, dto.Description, origin,
		dto.DefaultWorkCenterID, dto.StandardTime, dto.SetupTime, dto.CreatedBy)
	if err != nil {
		return nil, err
	}

	created, err := uc.repo.CreateOperation(ctx, op)
	if err != nil {
		return nil, err
	}
	return toOperationResponse(created), nil
}

func (uc *OperationUseCase) Update(ctx context.Context, dto request.UpdateOperationDTO) (*response.OperationResponse, error) {
	op, err := uc.repo.GetOperationByID(ctx, dto.ID)
	if err != nil {
		return nil, fmt.Errorf("operation not found: %w", err)
	}
	op.Name = dto.Name
	op.Description = dto.Description
	op.Origin = entity.OperationOrigin(dto.Origin)
	op.Situation = entity.OperationSituation(dto.Situation)
	op.DefaultWorkCenterID = dto.DefaultWorkCenterID
	op.StandardTime = dto.StandardTime
	op.SetupTime = dto.SetupTime

	updated, err := uc.repo.UpdateOperation(ctx, op)
	if err != nil {
		return nil, err
	}
	return toOperationResponse(updated), nil
}

func (uc *OperationUseCase) GetByID(ctx context.Context, id int64) (*response.OperationResponse, error) {
	op, err := uc.repo.GetOperationByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("operation not found: %w", err)
	}
	return toOperationResponse(op), nil
}

func (uc *OperationUseCase) List(ctx context.Context, onlyActive bool) ([]*response.OperationResponse, error) {
	ops, err := uc.repo.ListOperations(ctx, onlyActive)
	if err != nil {
		return nil, err
	}
	out := make([]*response.OperationResponse, 0, len(ops))
	for _, op := range ops {
		out = append(out, toOperationResponse(op))
	}
	return out, nil
}

func (uc *OperationUseCase) Deactivate(ctx context.Context, id int64) error {
	return uc.repo.DeactivateOperation(ctx, id)
}

func toOperationResponse(op *entity.Operation) *response.OperationResponse {
	return &response.OperationResponse{
		ID:                  op.ID,
		Code:                op.Code,
		Name:                op.Name,
		Description:         op.Description,
		Origin:              string(op.Origin),
		Situation:           string(op.Situation),
		DefaultWorkCenterID: op.DefaultWorkCenterID,
		StandardTime:        op.StandardTime,
		SetupTime:           op.SetupTime,
		IsActive:            op.IsActive,
		CreatedAt:           op.CreatedAt,
	}
}

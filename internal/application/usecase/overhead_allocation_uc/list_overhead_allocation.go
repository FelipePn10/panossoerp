package overhead_allocation_uc

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/overhead_allocation/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/overhead_allocation/repository"
)

type ListOverheadAllocationsUseCase struct {
	Repo repository.OverheadAllocationRepository
	Auth ports.AuthService
}

func (uc *ListOverheadAllocationsUseCase) Execute(
	ctx context.Context,
) ([]*entity.OverheadAllocation, error) {
	if !uc.Auth.CanListOverheadAllocation(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	return uc.Repo.List(ctx)
}

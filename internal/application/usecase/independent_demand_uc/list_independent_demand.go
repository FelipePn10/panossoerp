package independent_demand_uc

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/independent_demand/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/independent_demand/repository"
)

type ListIndependentDemandsUseCase struct {
	Repo repository.IndependentDemandRepository
	Auth ports.AuthService
}

func (uc *ListIndependentDemandsUseCase) Execute(
	ctx context.Context,
) ([]*entity.IndependentDemand, error) {
	if !uc.Auth.CanListIndependentDemand(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}
	return uc.Repo.List(ctx)
}

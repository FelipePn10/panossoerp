package independent_demand_uc

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/independent_demand/repository"
)

type DeleteIndependentDemandUseCase struct {
	Repo repository.IndependentDemandRepository
	Auth ports.AuthService
}

func (uc *DeleteIndependentDemandUseCase) Execute(
	ctx context.Context,
	code int64,
) error {
	if !uc.Auth.CanDeleteIndependentDemand(ctx) {
		return errorsuc.ErrUnauthorized
	}

	return uc.Repo.Delete(ctx, code)
}

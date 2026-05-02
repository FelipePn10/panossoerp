package independent_demand_uc

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/independent_demand/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/independent_demand/repository"
)

type ListIndependentDemandByItemUseCase struct {
	Repo repository.IndependentDemandRepository
	Auth ports.AuthService
}

func (uc *ListIndependentDemandByItemUseCase) Execute(
	ctx context.Context,
	itemCode int64,
) ([]*entity.IndependentDemand, error) {
	if !uc.Auth.CanViewIndependentDemand(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	return uc.Repo.ListByItem(ctx, itemCode)
}

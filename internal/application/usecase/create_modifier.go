package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/modifier/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/modifier/repository"
)

type CreateModifierUseCase struct {
	repo repository.ModifierRepository
	auth ports.AuthService
}

func (uc *CreateModifierUseCase) Execute(
	ctx context.Context,
	modifier *entity.Modifier,
) (*entity.Modifier, error) {
	if !uc.auth.CanCreateModifier(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	created, err := uc.repo.Create(ctx, modifier)
	if err != nil {
		return nil, err
	}

	return created, nil
}

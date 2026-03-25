package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/group/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/group/repository"
)

type CreateGroupUseCase struct {
	repo repository.GroupRepository
	auth ports.AuthService
}

func (uc *CreateGroupUseCase) Execute(
	ctx context.Context,
	group *entity.Group,
) (*entity.Group, error) {
	if !uc.auth.CanCreateGroup(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	created, err := uc.repo.Create(ctx, group)
	if err != nil {
		return nil, err
	}
	return created, nil
}

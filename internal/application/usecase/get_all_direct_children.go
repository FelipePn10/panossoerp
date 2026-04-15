package usecase

import (
	"context"
	"fmt"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/application/dto/response"
	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/structure/entity"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/repository/structure"
)

type ItemStructureRepository interface {
	GetAllDirectChildren(ctx context.Context, parentItemID int64) ([]*entity.ItemStructure, error)
}

type GetAllDirectChildrenUseCase struct {
	repo ItemStructureRepository
	auth ports.AuthService
}

func (uc *GetAllDirectChildrenUseCase) Execute(
	ctx context.Context,
	dto request.GetAllDirectChildrenDTO,
) ([]*response.StructureComponentResponse, error) {
	if !uc.auth.GetAllStructure(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	if dto.ParentItemID <= 0 {
		return nil, fmt.Errorf("parentItemId invalid")
	}

	items, err := uc.repo.GetAllDirectChildren(ctx, dto.ParentItemID)
	if err != nil {
		return nil, fmt.Errorf("error when searching for direct children: %w", err)
	}

	return structure.ToItemStructureListDTO(items), nil
}

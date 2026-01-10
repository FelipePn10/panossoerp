package usecase

import (
	"context"
	"errors"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/bom/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/bom/repository"
)

type CreateBomUseCase struct {
	repo repository.BomRepository
}

func (uc *CreateBomUseCase) Execute(
	ctx context.Context,
	dto request.CreateBomUseCaseRequestDTO,
) (*entity.Bom, error) {
	bom, err := entity.NewBom(
		dto.ProductId,
		dto.BomType,
		dto.Version,
		dto.ValidFrom,
		dto.Status,
	)
	if err != nil {
		if errors.Is(err, repository.ErrInvalidBom) {
			return nil, errorsuc.ErrCreateBom
		}
		return nil, err
	}
	return uc.repo.Create(ctx, bom)
}

package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/structure/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/structure/repository"
)

// UpdateStructureComponentUseCase atualiza quantidade, unidade de medida,
// percentual de perda, posição e notas de um componente BOM existente.
// Nota: a máscara (parent_mask) e os IDs pai/filho NÃO são editáveis.
// Para mudar esses campos, remova e recrie o componente.
type UpdateStructureComponentUseCase struct {
	repo repository.ItemStructureRepository
	auth ports.AuthService
}

func (uc *UpdateStructureComponentUseCase) Execute(
	ctx context.Context,
	code int64,
	dto request.UpdateStructureComponentDTO,
) (*entity.ItemStructure, error) {

	if !uc.auth.UpdateStructure(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	structure, err := uc.repo.GetByID(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("structural component %d not found: %w", code, err)
	}

	if !structure.IsActive {
		return nil, errors.New("it is not possible to update an inactive component")
	}

	if err := structure.Update(
		dto.Quantity,
		dto.UnitOfMeasurement,
		dto.Health,
		dto.LossPercentage,
		dto.Position,
		dto.Notes,
	); err != nil {
		return nil, err
	}

	return uc.repo.Update(ctx, structure)
}

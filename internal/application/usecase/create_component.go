package usecase

// import (
// 	"context"

// 	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
// 	"github.com/FelipePn10/panossoerp/internal/domain/component/entity"
// 	"github.com/FelipePn10/panossoerp/internal/domain/component/repository"
// 	"github.com/FelipePn10/panossoerp/internal/domain/component/valueobject"
// )

// type CreateComponentUseCase struct {
// 	repo repository.ComponentRepository
// }

// func (uc *CreateComponentUseCase) Execute(
// 	ctx context.Context,
// 	dto request.CreateComponentRequestDTO,
// ) (*entity.Component, error) {
// 	if dto.Warehouse < 0 {

// 	}

// 	code, err := valueobject.NewComponentCode(dto.GroupCode)
// 	if err != nil {
// 		return &entity.Component{}, err
// 	}

// 	component, err := entity.NewComponent(
// 		code.String(),
// 		dto.GroupCode,
// 		dto.Name,
// 		dto.Warehouse,
// 		dto.CreatedBy,
// 	)
// 	if err != nil {
// 		return &entity.Component{}, err
// 	}
// 	return uc.repo.Save(ctx, component)
// }

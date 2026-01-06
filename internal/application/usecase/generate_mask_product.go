package usecase

import (
	"context"
	"errors"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/domain/generate_mask_for_product/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/generate_mask_for_product/repository"
	"github.com/FelipePn10/panossoerp/internal/domain/generate_mask_for_product/valueobject"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/auth"
	contextkey "github.com/FelipePn10/panossoerp/internal/interfaces/http/context"
	"github.com/google/uuid"
)

type GenerateMaskForProductUseCase struct {
	repo repository.GenerateMaskForProductRepository
}

func (uc *GenerateMaskForProductUseCase) Execute(
	ctx context.Context,
	dto request.GenerateMaskProductRequestDTO,
) error {
	claims, ok := ctx.Value(contextkey.UserKey).(*auth.UserClaims)
	if !ok || claims.UserID == "" {
		return errors.New("unauthenticated user")
	}

	userUUID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return errors.New("invalid user id")
	}

	answers := make([]valueobject.MaskAnswer, 0, len(dto.Answers))

	for _, a := range dto.Answers {
		optionValue, err := uc.repo.GetOptionValue(ctx, a.OptionID)
		if err != nil {
			return err
		}

		answer, err := valueobject.NewMaskAnswer(
			a.QuestionID,
			a.OptionID,
			a.Position,
			optionValue,
		)
		if err != nil {
			return err
		}

		answers = append(answers, answer)
	}

	mask, err := valueobject.NewProductMask(dto.ProductCode, answers)
	if err != nil {
		return err
	}

	productMask := entity.ProductMask{
		ProductCode: dto.ProductCode,
		Mask:        mask.Value(),
		MaskHash:    mask.Hash(),
		CreatedBy:   userUUID,
	}

	return uc.repo.Generate(ctx, &productMask)
}

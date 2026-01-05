package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/domain/generate_mask_for_product/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/generate_mask_for_product/repository"
	"github.com/FelipePn10/panossoerp/internal/domain/generate_mask_for_product/valueobject"
)

type GenerateMaskForProduct struct {
	repo repository.GenerateMaskForProductRepository
}

func (uc *GenerateMaskForProduct) Execute(
	ctx context.Context,
	dto request.GenerateMaskProductRequestDTO,
) error {

	answers := make([]valueobject.MaskAnswer, 0, len(dto.Answers))

	for _, a := range dto.Answers {
		answer, err := valueobject.NewMaskAnswer(
			a.QuestionID,
			a.OptionID,
			a.Position,
		)
		if err != nil {
			return err
		}

		answers = append(answers, answer)
	}

	mask, err := valueobject.NewProductMask(answers)
	if err != nil {
		return err
	}

	productMask := entity.ProductMask{
		ProductCode: dto.ProductCode,
		Mask:        mask.Value(),
		MaskHash:    mask.Hash(),
		CreatedBy:   dto.CreatedBy,
	}

	return uc.repo.Generate(ctx, productMask)
}

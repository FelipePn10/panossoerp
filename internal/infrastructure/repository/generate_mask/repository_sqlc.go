package generatemask

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/generate_mask_for_product/entity"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

func (r *repositoryGenerateMaskSQLC) Generate(
	ctx context.Context,
	mask *entity.ProductMask,
) error {
	params := sqlc.InsertProductMaskParams{
		ProductCode: mask.ProductCode,
		Mask:        mask.Mask,
		MaskHash:    mask.MaskHash,
		CreatedBy:   mask.CreatedBy,
	}
	maskRecord, err := r.q.InsertProductMask(ctx, params)
	if err != nil {
		return err
	}

	for _, ans := range mask.Answers {
		err := r.q.InsertProductMaskAnswer(ctx, sqlc.InsertProductMaskAnswerParams{
			MaskID:     maskRecord.ID,
			QuestionID: ans.QuestionID(),
			OptionID:   ans.OptionID(),
			Position:   ans.Position(),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *repositoryGenerateMaskSQLC) GetOptionValue(ctx context.Context, optionID int64) (string, error) {
	value, err := r.q.GetOptionValueByID(ctx, optionID)
	if err != nil {
		return "", err
	}
	return value, nil
}

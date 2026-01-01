package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/domain/questions_options/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/questions_options/repository"
)

type CreateQuestionOptionUseCase struct {
	repo repository.QuestionsOptionsRepository
}

func (uc *CreateQuestionOptionUseCase) Execute(
	ctx context.Context,
	dto request.CreateQuestionOptionRequest,
) error {
	qstops, err := entity.NewQuestionsOptions(
		dto.Value,
		dto.QuestionId,
		dto.CreatedBy,
	)
	if err != nil {
		return err
	}
	return uc.repo.Save(ctx, qstops)
}

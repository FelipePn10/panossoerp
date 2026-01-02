package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/questions_options/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/questions_options/repository"
)

type DeleteQuestionOptionUseCase struct {
	repo repository.QuestionsOptionsRepository
}

func (uc *DeleteQuestionOptionUseCase) Execute(
	ctx context.Context,
	questionid int64,
) error {
	if err := entity.ValidateQuestionOptionDeletion(questionid); err != nil {
		return err
	}
	return uc.repo.Delete(ctx, questionid)
}

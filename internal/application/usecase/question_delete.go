package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/questions/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/questions/repository"
)

type DeleteQuestionUseCase struct {
	repo repository.QuestionsRepository
}

func (uc *DeleteQuestionUseCase) Execute(
	ctx context.Context,
	id int64,
) error {
	if err := entity.ValidateQuestionDeletion(id); err != nil {
		return err
	}
	return uc.repo.Delete(ctx, id)
}

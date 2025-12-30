package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/domain/questions/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/questions/repository"
)

type CreateQuestion struct {
	repo repository.QuestionsRepository
}

func (uc *CreateQuestion) Execute(
	ctx context.Context,
	dto request.CreateQuestionRequestDTO,
) error {
	qst, err := entity.NewQuestion(
		dto.Name,
		dto.CreatedBy,
	)
	if err != nil {
		return err
	}
	return uc.repo.Save(ctx, qst)
}

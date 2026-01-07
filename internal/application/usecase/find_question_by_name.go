package usecase

import (
	"context"
	"errors"

	"github.com/FelipePn10/panossoerp/internal/domain/questions/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/questions/repository"
)

type FindQuestionByName struct {
	repo repository.QuestionsRepository
}

func (uc *FindQuestionByName) Execute(
	ctx context.Context,
	name string,
) (*entity.Question, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	return uc.repo.FindQuestionByName(ctx, name)
}

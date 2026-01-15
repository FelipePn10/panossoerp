package usecase

import (
	"context"
	"errors"
	"strings"

	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
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
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errorsuc.ErrInvalidSearchParams
	}

	question, err := uc.repo.FindQuestionByName(ctx, name)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, errorsuc.ErrQuestionNotFound
		}
		return nil, err
	}
	return question, nil
}

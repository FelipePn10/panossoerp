package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/questions/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/questions/repository"
)

type CreateQuestion struct {
	repo repository.QuestionsRepository
	auth ports.AuthService
}

func (uc *CreateQuestion) Execute(
	ctx context.Context,
	dto request.CreateQuestionRequestDTO,
) (*entity.Question, error) {
	if !uc.auth.CanCreateQuestion(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	exists, err := uc.repo.ExistsByName(ctx, dto.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errorsuc.ErrQuestionAlreadyExists
	}

	qst, err := entity.NewQuestion(
		dto.Name,
		dto.CreatedBy,
	)
	if err != nil {
		return nil, err
	}

	crate, err := uc.repo.Save(ctx, qst)
	if err != nil {
		return nil, err
	}
	return crate, nil
}

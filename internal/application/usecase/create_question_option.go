package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/questions_options/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/questions_options/repository"
)

type CreateQuestionOptionUseCase struct {
	repo repository.QuestionsOptionsRepository
	auth ports.AuthService
}

func (uc *CreateQuestionOptionUseCase) Execute(
	ctx context.Context,
	dto request.CreateQuestionOptionRequest,
) (*entity.QuestionsOptions, error) {
	if !uc.auth.CanCreateQuestionOption(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	exists, err := uc.repo.ExistsByValue(ctx, dto.Value)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errorsuc.ErrQuestionOptionAlreadyExists
	}

	qstops, err := entity.NewQuestionsOptions(
		dto.Value,
		dto.QuestionId,
		dto.CreatedBy,
	)
	if err != nil {
		return nil, err
	}

	create, err := uc.repo.Save(ctx, qstops)
	if err != nil {
		return nil, err
	}
	return create, nil
}

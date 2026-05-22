package question_option_uc

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/questions_options/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/questions_options/repository"
)

type ListOptionsByQuestionUseCase struct {
	Repo repository.QuestionsOptionsRepository
	Auth ports.AuthService
}

func NewListOptionsByQuestionUseCase(
	repo repository.QuestionsOptionsRepository,
	auth ports.AuthService,
) *ListOptionsByQuestionUseCase {
	return &ListOptionsByQuestionUseCase{Repo: repo, Auth: auth}
}

func (uc *ListOptionsByQuestionUseCase) Execute(
	ctx context.Context,
	questionID int64,
) ([]entity.QuestionsOptions, error) {
	if !uc.Auth.CanCreateQuestionOption(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}
	return uc.Repo.ListByQuestionID(ctx, questionID)
}

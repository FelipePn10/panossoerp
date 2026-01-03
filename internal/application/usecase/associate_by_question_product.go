package usecase

import (
	"context"
	"errors"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/domain/associate_questions/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/associate_questions/repository"
)

var (
	ErrQuestionAlreadyLinked = errors.New("question already linked to product")
	ErrPositionAlreadyUsed   = errors.New("position already used for product")
)

type AssociateByQuestionProductUseCase struct {
	repo repository.ProductQuestionsRepository
}

func (uc *AssociateByQuestionProductUseCase) Execute(
	ctx context.Context,
	dto request.AssociateByQuestionProductRequestDTO,
) error {
	exists, err := uc.repo.ExistsByProductAndQuestion(
		ctx,
		dto.ProductID,
		dto.QuestionID,
	)
	if err != nil {
		return err
	}
	if exists {
		return ErrQuestionAlreadyLinked
	}

	positionUsed, err := uc.repo.ExistsByProductAndPostion(
		ctx,
		dto.ProductID,
		dto.QuestionID,
	)
	if err != nil {
		return err
	}
	if positionUsed {
		return ErrPositionAlreadyUsed
	}

	pq, err := entity.New(
		dto.ProductID,
		dto.QuestionID,
		dto.Position,
	)
	if err != nil {
		return err
	}

	return uc.repo.Create(ctx, pq)
}

package question_uc

import (
	"context"
	"errors"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/associate_questions/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/associate_questions/repository"
	itemrepo "github.com/FelipePn10/panossoerp/internal/domain/items/repository"
	"github.com/FelipePn10/panossoerp/internal/domain/items/valueobject"
)

var (
	ErrQuestionAlreadyLinked = errors.New("question already linked to product")
	ErrPositionAlreadyUsed   = errors.New("position already used for product")
)

type AssociateByQuestionItemUseCase struct {
	Repo     repository.AssociateQuestionsRepository
	ItemRepo itemrepo.ItemRepository
	Auth     ports.AuthService
}

func NewAssociateByQuestionItemUseCase(
	repo repository.AssociateQuestionsRepository,
	itemRepo itemrepo.ItemRepository,
	auth ports.AuthService,
) *AssociateByQuestionItemUseCase {
	return &AssociateByQuestionItemUseCase{
		Repo:     repo,
		ItemRepo: itemRepo,
		Auth:     auth,
	}
}

func (uc *AssociateByQuestionItemUseCase) Execute(
	ctx context.Context,
	dto request.AssociateByQuestionItemRequestDTO,
) error {
	if !uc.Auth.CanAssociateByQuestionProduct(ctx) {
		return errorsuc.ErrUnauthorized
	}

	itemCode, err := valueobject.NewItemCode(dto.ItemCode)
	if err != nil {
		return err
	}

	item, err := uc.ItemRepo.FindItemByCode(ctx, itemCode)
	if err != nil {
		return err
	}

	exists, err := uc.Repo.ExistsByItemAndQuestion(ctx, item.ID, dto.QuestionID)
	if err != nil {
		return err
	}
	if exists {
		return ErrQuestionAlreadyLinked
	}

	positionUsed, err := uc.Repo.ExistsByItemAndPosition(ctx, item.ID, dto.Position)
	if err != nil {
		return err
	}
	if positionUsed {
		return ErrPositionAlreadyUsed
	}

	pq, err := entity.New(item.ID, dto.QuestionID, dto.Position)
	if err != nil {
		return err
	}

	return uc.Repo.Associate(ctx, pq)
}

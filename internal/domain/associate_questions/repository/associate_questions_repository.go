package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/associate_questions/entity"
)

type AssociateQuestionsRepository interface {
	Associate(ctx context.Context, pq *entity.AssociateQuestion) error
	ExistsByItemAndQuestion(
		ctx context.Context,
		itemCode int64,
		questionID int64,
	) (bool, error)
	ExistsByItemAndPosition(
		ctx context.Context,
		itemID int64,
		position int,
	) (bool, error)
	GetByItemCode(ctx context.Context, itemCode int64) ([]entity.AssociateQuestionDetail, error)
	ListAll(ctx context.Context) ([]entity.ItemQuestionRow, error)
}

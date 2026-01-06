package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/associate_questions/entity"
)

type AssociateQuestionsRepository interface {
	Associate(ctx context.Context, pq *entity.AssociateQuestion) error
	ExistsByProductAndQuestion(
		ctx context.Context,
		ProductID int64,
		questionID int64,
	) (bool, error)
	ExistsByProductAndPosition(
		ctx context.Context,
		productID int64,
		position int,
	) (bool, error)
}

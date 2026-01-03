package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/associate_questions/entity"
)

type ProductQuestionsRepository interface {
	Create(ctx context.Context, pq *entity.ProductQuestion) error
	ExistsByProductAndQuestion(
		ctx context.Context,
		productID int64,
		questionID int64,
	) (bool, error)
	ExistsByProductAndPostion(
		ctx context.Context,
		productID int64,
		questionID int64,
	) (bool, error)
}

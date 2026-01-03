package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/product_questions/entity"
)

type ProductQuestionsRepository interface {
	Create(ctx context.Context, pq *entity.ProductQuestion)
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

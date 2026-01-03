package productquestion

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/associate_questions/entity"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

type ProductQuestionRepository struct {
	q *sqlc.Queries
}

func (r *ProductQuestionRepository) Create(
	ctx context.Context,
	pq *entity.ProductQuestion,
) error {
	return r.q.CreateProductQuestion(ctx, sqlc.CreateProductQuestionParams{
		ProductID:  pq.ProductID,
		QuestionID: pq.QuestionID,
		Position:   int32(pq.Position),
		CreatedAt:  pq.CreatedAt,
	})
}

func (r *ProductQuestionRepository) ExistsByProductAndQuestion(
	ctx context.Context,
	productID int64,
	questionID int64,
) (bool, error) {

	return r.q.ExistsProductQuestionByProductAndQuestion(
		ctx,
		sqlc.ExistsProductQuestionByProductAndQuestionParams{
			ProductID:  productID,
			QuestionID: questionID,
		},
	)
}

func (r *ProductQuestionRepository) ExistsByProductAndPosition(
	ctx context.Context,
	productID int64,
	position int,
) (bool, error) {

	return r.q.ExistsProductQuestionByProductAndPosition(
		ctx,
		sqlc.ExistsProductQuestionByProductAndPositionParams{
			ProductID: productID,
			Position:  int32(position),
		},
	)
}

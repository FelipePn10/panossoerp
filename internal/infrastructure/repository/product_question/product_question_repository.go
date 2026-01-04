package productquestion

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/associate_questions/entity"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

func (r *AssociateQuestionProductRepository) Associate(
	ctx context.Context,
	pq *entity.AssociateQuestion,
) error {
	return r.q.AssociateQuestionProduct(ctx, sqlc.AssociateQuestionProductParams{
		ProductID:  pq.ProductID,
		QuestionID: pq.QuestionID,
		Position:   int32(pq.Position),
		CreatedAt:  pq.CreatedAt,
	})
}

func (r *AssociateQuestionProductRepository) ExistsByProductAndQuestion(
	ctx context.Context,
	productID int64,
	questionID int64,
) (bool, error) {
	return r.q.ExistsByProductAndQuestion(ctx, sqlc.ExistsByProductAndQuestionParams{
		ProductID:  productID,
		QuestionID: questionID,
	})
}

func (r *AssociateQuestionProductRepository) ExistsByProductAndPosition(
	ctx context.Context,
	productID int64,
	position int,
) (bool, error) {
	return r.q.ExistsByProductAndPosition(ctx, sqlc.ExistsByProductAndPositionParams{
		ProductID: productID,
		Position:  int32(position),
	})
}

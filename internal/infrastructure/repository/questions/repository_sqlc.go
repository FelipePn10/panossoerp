package questions

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/questions/entity"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

func (r *repositoryQuestionSQLC) Save(
	ctx context.Context,
	qst *entity.Question,
) error {
	_, err := r.q.CreateQuestion(ctx, sqlc.CreateQuestionParams{
		Name:      qst.Name,
		Createdby: qst.CreatedBy,
	})
	return err
}

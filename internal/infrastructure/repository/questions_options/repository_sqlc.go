package questionsoptions

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/questions_options/entity"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

func (r *repositoryQuestionOptionsSQLC) Save(
	ctx context.Context,
	qstops *entity.QuestionsOptions,
) error {
	_, err := r.q.CreateQuestionOption(ctx, sqlc.CreateQuestionOptionParams{
		Value:      qstops.Value,
		Createdby:  qstops.CreatedBy,
		QuestionID: qstops.QuestionId,
	})
	return err
}

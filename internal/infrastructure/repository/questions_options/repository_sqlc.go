package questionsoptions

import (
	"context"
	"database/sql"
	"errors"

	"github.com/FelipePn10/panossoerp/internal/domain/product/repository"
	"github.com/FelipePn10/panossoerp/internal/domain/questions_options/entity"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

func (r *repositoryQuestionOptionsSQLC) Save(
	ctx context.Context,
	qstops *entity.QuestionsOptions,
) error {
	_, err := r.q.CreateQuestionOption(ctx, sqlc.CreateQuestionOptionParams{
		Value:      qstops.Value,
		CreatedBy:  qstops.CreatedBy,
		QuestionID: qstops.QuestionId,
	})
	return err
}

func (r *repositoryQuestionOptionsSQLC) ExistsQuestionOptionByValue(
	ctx context.Context,
	value string,
) (*entity.QuestionsOptions, error) {
	dbQuestionOption, err := r.q.ExistsQuestionOptionByValue(ctx, value)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
	}
	return &entity.QuestionsOptions{
		QuestionId: dbQuestionOption.ID,
		CreatedBy:  dbQuestionOption.CreatedBy,
		Value:      dbQuestionOption.Value,
	}, nil
}

func (r *repositoryQuestionOptionsSQLC) Delete(
	ctx context.Context,
	questionid int64,
) error {
	return r.q.DeleteQuestionOption(ctx, questionid)
}

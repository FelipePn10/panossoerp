package questionsoptions

import (
	"context"
	"database/sql"
	"errors"

	"github.com/FelipePn10/panossoerp/internal/domain/questions_options/entity"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

func (r *repositoryQuestionOptionsSQLC) Save(
	ctx context.Context,
	qstops *entity.QuestionsOptions,
) (*entity.QuestionsOptions, error) {
	dbQuestionOption, err := r.q.CreateQuestionOption(ctx, sqlc.CreateQuestionOptionParams{
		Value:      qstops.Value,
		CreatedBy:  qstops.CreatedBy,
		QuestionID: qstops.QuestionId,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return &entity.QuestionsOptions{
		QuestionId: dbQuestionOption.ID,
		CreatedBy:  dbQuestionOption.CreatedBy,
		Value:      dbQuestionOption.Value,
	}, nil
}

func (r *repositoryQuestionOptionsSQLC) ExistsQuestionOptionByValue(
	ctx context.Context,
	value string,
) (bool, error) {
	_, err := r.q.ExistsQuestionOptionByValue(ctx, value)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, err
		}
		return false, err
	}
	return true, nil
}

func (r *repositoryQuestionOptionsSQLC) Delete(
	ctx context.Context,
	questionid int64,
) error {
	return r.q.DeleteQuestionOption(ctx, questionid)
}

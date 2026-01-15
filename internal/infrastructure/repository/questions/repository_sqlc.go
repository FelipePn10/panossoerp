package questions

import (
	"context"
	"database/sql"
	"errors"

	"github.com/FelipePn10/panossoerp/internal/domain/questions/entity"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

func (r *repositoryQuestionSQLC) Save(
	ctx context.Context,
	qst *entity.Question,
) (*entity.Question, error) {
	params := sqlc.CreateQuestionParams{
		Name:      qst.Name,
		Createdby: qst.CreatedBy,
	}
	dbQuestion, err := r.q.CreateQuestion(ctx, params)
	if err != nil {
		return nil, err
	}

	return &entity.Question{
		Name:      dbQuestion.Name,
		CreatedBy: dbQuestion.Createdby,
	}, nil
}

func (r *repositoryQuestionSQLC) Delete(
	ctx context.Context,
	id int64,
) error {
	return r.q.DeleteQuestion(ctx, id)
}

func (r *repositoryQuestionSQLC) FindQuestionByName(
	ctx context.Context,
	name string,
) (*entity.Question, error) {
	dbQuestion, err := r.q.FindQuestionByName(ctx, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}
	return &entity.Question{
		Name:      dbQuestion.Name,
		CreatedBy: dbQuestion.Createdby,
	}, nil
}

func (r *repositoryQuestionSQLC) ExistsQuestionByName(
	ctx context.Context,
	name string,
) (bool, error) {
	_, err := r.q.ExistsQuestionByName(ctx, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, err
		}
		return false, err
	}
	return true, nil
}

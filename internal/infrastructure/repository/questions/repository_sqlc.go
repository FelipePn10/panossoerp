package questions

import (
	"context"
	"database/sql"
	"errors"

	"github.com/FelipePn10/panossoerp/internal/domain/questions/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/questions/repository"
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
	dbQuestion, err := r.q.FindQuestionByNameAndCode(ctx, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &entity.Question{
		Name:      dbQuestion.Name,
		CreatedBy: dbQuestion.Createdby,
	}, nil
}

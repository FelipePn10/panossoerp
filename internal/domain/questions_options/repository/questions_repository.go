package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/questions_options/entity"
)

type QuestionsOptionsRepository interface {
	Save(ctx context.Context, qstops *entity.QuestionsOptions) (*entity.QuestionsOptions, error)
	Delete(ctx context.Context, questionid int64) error
	ExistsQuestionOptionByValue(ctx context.Context, value string) (bool, error)
}

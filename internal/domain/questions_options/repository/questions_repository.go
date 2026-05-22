package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/questions_options/entity"
)

type QuestionsOptionsRepository interface {
	Save(ctx context.Context, qstops *entity.QuestionsOptions) (*entity.QuestionsOptions, error)
	Delete(ctx context.Context, questionid int64) error
	ExistsQuestionOptionByValue(ctx context.Context, value string, question_id int64) (bool, error)
	ListByQuestionID(ctx context.Context, questionID int64) ([]entity.QuestionsOptions, error)
}

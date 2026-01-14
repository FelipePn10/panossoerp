package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/questions/entity"
)

type QuestionsRepository interface {
	Save(ctx context.Context, qst *entity.Question) (*entity.Question, error)
	Delete(ctx context.Context, id int64) error
	FindQuestionByName(ctx context.Context, name string) (*entity.Question, error)
	ExistsByName(ctx context.Context, name string) (bool, error)
	//FindByID(ctx context.Context, id int64) (*entity.Question, error)
}

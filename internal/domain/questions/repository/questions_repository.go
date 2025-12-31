package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/questions/entity"
)

type QuestionsRepository interface {
	Save(ctx context.Context, qst *entity.Question) error
	Delete(ctx context.Context, id int64) error
	//FindByID(ctx context.Context, id int64) (*entity.Question, error)
}

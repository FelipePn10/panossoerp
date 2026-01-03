package productquestion

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

func NewProductQuestionRepositorySQLC(
	q *sqlc.Queries,
) *ProductQuestionRepository {
	return &ProductQuestionRepository{q: q}
}

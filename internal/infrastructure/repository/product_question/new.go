package productquestion

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

type AssociateQuestionProductRepository struct {
	q *sqlc.Queries
}

func NewAssociateQuestionProductRepositorySQLC(
	q *sqlc.Queries,
) *AssociateQuestionProductRepository {
	return &AssociateQuestionProductRepository{q: q}
}

package warehouse

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

type repositoryWarehouseSQLC struct {
	q *sqlc.Queries
}

func NewRepositoryQuestionSQLC(q *sqlc.Queries) *repositoryWarehouseSQLC {
	return &repositoryWarehouseSQLC{
		q: q,
	}
}

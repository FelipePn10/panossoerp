package product

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

type repositoryProductSQLC struct {
	q *sqlc.Queries
}

func NewRepositoryProductSQLC(q *sqlc.Queries) *repositoryProductSQLC {
	return &repositoryProductSQLC{
		q: q,
	}
}

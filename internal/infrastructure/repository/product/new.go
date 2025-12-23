package product

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

func NewRepositorySQLC(q *sqlc.Queries) *repositoryProductSQLC {
	return &repositoryProductSQLC{
		q: q,
	}
}

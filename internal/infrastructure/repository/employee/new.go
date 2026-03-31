package employee

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

type repositoryEmployeeSQLC struct {
	q *sqlc.Queries
}

func NewRepositoryEmployeeSQLC(q *sqlc.Queries) *repositoryEmployeeSQLC {
	return &repositoryEmployeeSQLC{
		q: q,
	}
}

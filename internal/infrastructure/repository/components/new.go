package components

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

type repositoryComponentsSQLC struct {
	q *sqlc.Queries
}

func NewRepositoryComponentsSQLC(q *sqlc.Queries) *repositoryComponentsSQLC {
	return &repositoryComponentsSQLC{
		q: q,
	}
}

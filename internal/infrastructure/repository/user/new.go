package user

import (
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

type repositoryUserSQLC struct {
	q *sqlc.Queries
}

func NewRepositoryUserSQLC(q *sqlc.Queries) *repositoryUserSQLC {
	return &repositoryUserSQLC{
		q: q,
	}
}

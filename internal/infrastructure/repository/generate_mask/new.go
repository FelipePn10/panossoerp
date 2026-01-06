package generatemask

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

type repositoryGenerateMaskSQLC struct {
	q *sqlc.Queries
}

func NewRepositoryGenerateMaskSQLC(q *sqlc.Queries) *repositoryGenerateMaskSQLC {
	return &repositoryGenerateMaskSQLC{
		q: q,
	}
}

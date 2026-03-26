package enterprise

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

type repositoryEnterpriseSQLC struct {
	q *sqlc.Queries
}

func NewRepositoryEnterpriseSQLC(q *sqlc.Queries) *repositoryEnterpriseSQLC {
	return &repositoryEnterpriseSQLC{
		q: q,
	}
}

package group

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

type repositoryGroupSQLC struct {
	q *sqlc.Queries
}

func NewRepositoryGroupSQLC(q *sqlc.Queries) *repositoryGroupSQLC {
	return &repositoryGroupSQLC{
		q: q,
	}
}

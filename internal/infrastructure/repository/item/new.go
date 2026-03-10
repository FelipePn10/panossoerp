package item

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

type repositoryItemSQLC struct {
	q *sqlc.Queries
}

func NewRepositoryItemSQLC(q *sqlc.Queries) *repositoryItemSQLC {
	return &repositoryItemSQLC{
		q: q,
	}
}

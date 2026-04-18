package item

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

type RepositoryItemSQLC struct {
	q *sqlc.Queries
}

func NewRepositoryItemSQLC(q *sqlc.Queries) *RepositoryItemSQLC {
	return &RepositoryItemSQLC{
		q: q,
	}
}

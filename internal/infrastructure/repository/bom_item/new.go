package bomitem

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

type repositoryBomItemSQLC struct {
	q *sqlc.Queries
}

func NewRepositoryBomItemSQLC(q *sqlc.Queries) *repositoryBomItemSQLC {
	return &repositoryBomItemSQLC{
		q: q,
	}
}

package modifier

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

type repositoryModifierSQLC struct {
	q *sqlc.Queries
}

func NewRepositoryModifierSQLC(q *sqlc.Queries) *repositoryModifierSQLC {
	return &repositoryModifierSQLC{
		q: q,
	}
}

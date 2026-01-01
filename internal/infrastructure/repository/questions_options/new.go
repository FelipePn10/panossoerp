package questionsoptions

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

type repositoryQuestionOptionsSQLC struct {
	q *sqlc.Queries
}

func NewRepositoryQuestionOptionSQLC(q *sqlc.Queries) *repositoryQuestionOptionsSQLC {
	return &repositoryQuestionOptionsSQLC{
		q: q,
	}
}

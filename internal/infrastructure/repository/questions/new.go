package questions

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

type repositoryQuestionSQLC struct {
	q *sqlc.Queries
}

func NewRepositoryQuestionSQLC(q *sqlc.Queries) *repositoryQuestionSQLC {
	return &repositoryQuestionSQLC{
		q: q,
	}
}

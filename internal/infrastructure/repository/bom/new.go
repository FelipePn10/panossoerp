package bom

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

type repositoryBomSQLC struct {
	q *sqlc.Queries
}

func NewRepostioryBomSQLC(q *sqlc.Queries) *repositoryBomSQLC {
	return &repositoryBomSQLC{
		q: q,
	}
}

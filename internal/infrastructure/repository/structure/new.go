package structure

import (
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

type ItemStructureRepositorySQLC struct {
	q *sqlc.Queries
}

func NewItemStructureRepository(q *sqlc.Queries) *ItemStructureRepositorySQLC {
	return &ItemStructureRepositorySQLC{
		q: q,
	}
}

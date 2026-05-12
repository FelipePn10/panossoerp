package overhead_allocation

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

type OverheadAllocationRepositorySQLC struct {
	q *sqlc.Queries
}

func NewOverheadAllocationRepositorySQLC(q *sqlc.Queries) *OverheadAllocationRepositorySQLC {
	return &OverheadAllocationRepositorySQLC{q: q}
}

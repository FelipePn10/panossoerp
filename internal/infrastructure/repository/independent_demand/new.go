package independent_demand

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

type IndependentDemandRepositorySQLC struct {
	q *sqlc.Queries
}

func NewIndependentDemandRepositorySQLC(q *sqlc.Queries) *IndependentDemandRepositorySQLC {
	return &IndependentDemandRepositorySQLC{q: q}
}

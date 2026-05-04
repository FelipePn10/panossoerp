package item_calendar_promise

import "github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"

type ItemCalendarPromiseRepositorySQLC struct {
	q *sqlc.Queries
}

func NewItemCalendarPromiseRepositorySQLC(q *sqlc.Queries) *ItemCalendarPromiseRepositorySQLC {
	return &ItemCalendarPromiseRepositorySQLC{q: q}
}

package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/item_calendar_promise/entity"
)

type ItemCalendarPromiseRepository interface {
	UpsertDay(ctx context.Context, c *entity.ItemCalendarPromise) (*entity.ItemCalendarPromise, error)
	GetDay(ctx context.Context, itemCode int64, mask string, year, month, day int) (*entity.ItemCalendarPromise, error)
	GetWorkdaysInMonth(ctx context.Context, itemCode int64, mask string, year, month int) ([]*entity.ItemCalendarPromise, error)
	ListMonth(ctx context.Context, itemCode int64, mask string, year, month int) ([]*entity.ItemCalendarPromise, error)
	DeleteDay(ctx context.Context, itemCode int64, mask string, year, month, day int) error
}

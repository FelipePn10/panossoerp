package item_calendar_promise_uc

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/item_calendar_promise/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/item_calendar_promise/repository"
)

type ManageItemCalendarPromiseUseCase struct {
	Repo repository.ItemCalendarPromiseRepository
	Auth ports.AuthService
}

func (uc *ManageItemCalendarPromiseUseCase) UpsertDay(
	ctx context.Context,
	dto request.CreateItemCalendarDayDTO,
) (*entity.ItemCalendarPromise, error) {
	if !uc.Auth.CanManageItemCalendarPromise(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	cal := &entity.ItemCalendarPromise{
		ItemCode:    dto.ItemCode,
		Mask:        dto.Mask,
		Year:        dto.Year,
		Month:       dto.Month,
		Day:         dto.Day,
		IsWorkday:   dto.IsWorkday,
		Description: dto.Description,
	}

	return uc.Repo.UpsertDay(ctx, cal)
}

func (uc *ManageItemCalendarPromiseUseCase) GetDay(
	ctx context.Context,
	itemCode int64,
	mask string,
	year, month, day int,
) (*entity.ItemCalendarPromise, error) {
	if !uc.Auth.CanManageItemCalendarPromise(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}
	return uc.Repo.GetDay(ctx, itemCode, mask, year, month, day)
}

func (uc *ManageItemCalendarPromiseUseCase) ListMonth(
	ctx context.Context,
	itemCode int64,
	mask string,
	year, month int,
) ([]*entity.ItemCalendarPromise, error) {
	if !uc.Auth.CanManageItemCalendarPromise(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}
	return uc.Repo.ListMonth(ctx, itemCode, mask, year, month)
}

func (uc *ManageItemCalendarPromiseUseCase) GetWorkdaysInMonth(
	ctx context.Context,
	itemCode int64,
	mask string,
	year, month int,
) ([]*entity.ItemCalendarPromise, error) {
	if !uc.Auth.CanManageItemCalendarPromise(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}
	return uc.Repo.GetWorkdaysInMonth(ctx, itemCode, mask, year, month)
}

func (uc *ManageItemCalendarPromiseUseCase) DeleteDay(
	ctx context.Context,
	itemCode int64,
	mask string,
	year, month, day int,
) error {

	if !uc.Auth.CanManageItemCalendarPromise(ctx) {
		return errorsuc.ErrUnauthorized
	}

	return uc.Repo.DeleteDay(ctx, itemCode, mask, year, month, day)
}

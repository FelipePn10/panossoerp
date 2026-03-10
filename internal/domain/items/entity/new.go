package entity

import (
	"errors"

	"github.com/FelipePn10/panossoerp/internal/domain/enums/types"
	"github.com/google/uuid"
)

var (
	ErrInvalidWarehouseID = errors.New("warehouse_id cannot be empty")
	ErrInvalidName        = errors.New("name cannot be empty")
	ErrInvalidDesciption  = errors.New("description cannot be empty")
	ErrInvalidUserId      = errors.New("user_id cannot be empty")
)

func NewItem(
	warehouse_id int64,
	code string,
	name string,
	description string,
	types types.Type,
	status types.Status,
	health types.Health,
	created_by uuid.UUID,
) (*Item, error) {
	switch {
	case warehouse_id < 0:
		return nil, ErrInvalidWarehouseID
	case name == "":
		return nil, ErrInvalidName
	case description == "":
		return nil, ErrInvalidDesciption
	case created_by == uuid.Nil:
		return nil, ErrInvalidUserId
	}

	return &Item{
		WarehouseID: warehouse_id,
		Code:        code,
		Name:        name,
		Description: description,
		Type:        types,
		Status:      status,
		Health:      health,
		CreatedBy:   created_by,
	}, nil
}

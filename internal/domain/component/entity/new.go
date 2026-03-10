package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidCode      = errors.New("code cannot be empty")
	ErrInvalidName      = errors.New("name cannot be empty")
	ErrInvalidWarehouse = errors.New("ware cannot be empty")
	ErrInvalidGroupCode = errors.New("groupCode must be greater than zero")
)

func NewComponent(
	name string,
	group_code string,
	code string,
	warehouse int64,
	created_by uuid.UUID,
) (*Component, error) {
	switch {
	case name == "":
		return nil, ErrInvalidName
	case group_code == "":
		return nil, ErrInvalidGroupCode
	case code == "":
		return nil, ErrInvalidCode
	case warehouse < 0:
		return nil, ErrInvalidWarehouse
	}

	return &Component{
		Name:      name,
		GroupCode: group_code,
		Code:      code,
		Warehouse: warehouse,
		CreatedBy: created_by,
	}, nil
}

func ValidateComponentDeletion(id int64) error {
	if id < 0 {
		return errors.New("product id must be greater than zero")
	}
	return nil
}

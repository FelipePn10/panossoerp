package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidCode      = errors.New("code cannot be empty")
	ErrInvalidName      = errors.New("name cannot be empty")
	ErrInvalidGroupCode = errors.New("groupCode must be greater than zero")
)

func NewProduct(
	id uuid.UUID,
	code string,
	groupCode int16,
	name string,
	createdBy uuid.UUID,
) (*Product, error) {

	switch {
	case code == "":
		return nil, ErrInvalidCode
	case name == "":
		return nil, ErrInvalidName
	case groupCode <= 0:
		return nil, ErrInvalidGroupCode
	case createdBy == uuid.Nil:
		return nil, errors.New("createdBy cannot be nil UUID")
	}

	return &Product{
		ID:        id,
		Code:      code,
		GroupCode: groupCode,
		Name:      name,
		CreatedBy: createdBy,
	}, nil
}

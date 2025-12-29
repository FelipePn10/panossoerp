package entity

import (
	"errors"

	"github.com/FelipePn10/panossoerp/internal/domain/product/valueobject"
	"github.com/google/uuid"
)

var (
	ErrInvalidCode      = errors.New("code cannot be empty")
	ErrInvalidName      = errors.New("name cannot be empty")
	ErrInvalidGroupCode = errors.New("groupCode must be greater than zero")
)

func NewProduct(
	code valueobject.ProductCode,
	groupCode string,
	name string,
	createdBy uuid.UUID,
) (*Product, error) {

	switch {
	case code == valueobject.ProductCode{}:
		return nil, ErrInvalidCode
	case name == "":
		return nil, ErrInvalidName
	case groupCode <= "":
		return nil, ErrInvalidGroupCode
	case createdBy == uuid.Nil:
		return nil, errors.New("createdBy cannot be nil UUID")
	}

	return &Product{
		ID:        uuid.New(),
		Code:      code,
		GroupCode: groupCode,
		Name:      name,
		CreatedBy: createdBy,
	}, nil
}

func ValidateProductDeletion(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("id cannot be nil UUID")
	}
	return nil
}

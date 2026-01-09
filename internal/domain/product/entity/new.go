package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidCode      = errors.New("code cannot be empty")
	ErrInvalidName      = errors.New("name cannot be empty")
	ErrInvalidGroupCode = errors.New("groupCode must be greater than zero")
)

func NewProduct(
	code string,
	group_code string,
	name string,
	createdBy uuid.UUID,
	uom string,
) (*Product, error) {

	switch {
	case code == "":
		return nil, ErrInvalidCode
	case name == "":
		return nil, ErrInvalidName
	case group_code <= "":
		return nil, ErrInvalidGroupCode
	case createdBy == uuid.Nil:
		return nil, errors.New("createdBy cannot be nil UUID")
	}

	id := int64(time.Now().UnixNano())
	return &Product{
		ID:        id,
		Code:      code,
		GroupCode: group_code,
		Name:      name,
		UOM:       uom,
		CreatedBy: createdBy,
	}, nil
}

func ValidateProductDeletion(id int64) error {
	if id == 0 {
		return errors.New("product id must be greater than zero")
	}
	return nil
}

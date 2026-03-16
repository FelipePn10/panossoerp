package entity

import (
	"errors"

	"github.com/google/uuid"
)

func NewWarehouse(
	name string,
	description string,
	code string,
	types string,
	created_by uuid.UUID,
) (*Warehouse, error) {
	switch {
	case name == "":
		return nil, errors.ErrUnsupported
	case code == "":
		return nil, errors.ErrUnsupported
	case types == "":
		return nil, errors.ErrUnsupported
	case created_by == uuid.Nil:
		return nil, errors.New("createdby cannot be nil UUID")
	}
	return &Warehouse{
		Name:        name,
		Description: description,
		Code:        code,
		Type:        types,
		CreatedBy:   created_by,
	}, nil
}

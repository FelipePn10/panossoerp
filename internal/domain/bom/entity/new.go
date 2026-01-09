package entity

import (
	"errors"
	"time"
)

var (
	ErrInvalidProductId = errors.New("product_id must be greater than zero")
	ErrInvalidBomType   = errors.New("bom_type cannot be empty")
	ErrInvalidVersion   = errors.New("version must be greater than zero")
	ErrInvalidStatus    = errors.New("provide the correct status for the system to generate the BOM.")
)

func NewBom(
	product_id int64,
	bom_type string,
	version int,
	valid_from time.Time,
	status string,
) (*Bom, error) {
	switch {
	case product_id <= 0:
		return nil, ErrInvalidProductId
	case bom_type == "":
		return nil, ErrInvalidBomType
	case version <= 0:
		return nil, ErrInvalidVersion
	case status == "":
		return nil, ErrInvalidStatus
	}

	return &Bom{
		ProductId: product_id,
		BomType:   bom_type,
		Version:   version,
		ValidFrom: valid_from,
		Status:    status,
	}, nil
}

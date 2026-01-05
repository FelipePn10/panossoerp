package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidProductCode = errors.New("productCode cannot be empty")
	ErrInvalidMask        = errors.New("mask cannot be empty")
	ErrInvalidMaskHash    = errors.New("maskHash cannot be empty")
)

func NewProductMask(
	productCode string,
	mask string,
	maskHash string,
	createdBy uuid.UUID,
) (*ProductMask, error) {
	switch {
	case productCode == "":
		return nil, ErrInvalidProductCode
	case mask == "":
		return nil, ErrInvalidMask
	case maskHash == "":
		return nil, ErrInvalidMaskHash
	case createdBy == uuid.Nil:
		return nil, errors.New("createdBy cannot be nil UUID")
	}
	return &ProductMask{
		ProductCode: productCode,
		Mask:        mask,
		MaskHash:    maskHash,
		CreatedBy:   createdBy,
	}, nil
}

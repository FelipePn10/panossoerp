package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidItemCode = errors.New("item code cannot be empty")
	ErrInvalidMask     = errors.New("mask cannot be empty")
	ErrInvalidMaskHash = errors.New("maskHash cannot be empty")
)

func NewItemMask(
	itemCode string,
	mask string,
	maskHash string,
	createdBy uuid.UUID,
) (*ItemMask, error) {
	switch {
	case itemCode == "":
		return nil, ErrInvalidItemCode
	case mask == "":
		return nil, ErrInvalidMask
	case maskHash == "":
		return nil, ErrInvalidMaskHash
	case createdBy == uuid.Nil:
		return nil, errors.New("createdBy cannot be nil UUID")
	}
	return &ItemMask{
		ItemCode:  itemCode,
		Mask:      mask,
		MaskHash:  maskHash,
		CreatedBy: createdBy,
	}, nil
}

package entity

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidItemCode = errors.New("item code cannot be empty")
	ErrInvalidMask     = errors.New("mask cannot be empty")
	ErrInvalidMaskHash = errors.New("maskHash cannot be empty")
	ErrInvalidHash     = errors.New("maskHash does not match mask")
	ErrInvalidUser     = errors.New("createdBy cannot be nil UUID")
)

func NewItemMask(
	itemCode string,
	itemID int64,
	mask string,
	maskHash string,
	createdBy uuid.UUID,
) (*ItemMask, error) {

	// Validação estrutural
	if itemCode == "" {
		return nil, ErrInvalidItemCode
	}

	if mask == "" {
		return nil, ErrInvalidMask
	}

	if maskHash == "" {
		return nil, ErrInvalidMaskHash
	}

	if createdBy == uuid.Nil {
		return nil, ErrInvalidUser
	}

	expectedHash := generateHash(mask)
	if maskHash != expectedHash {
		return nil, ErrInvalidHash
	}

	return &ItemMask{
		ItemCode:  itemCode,
		ItemID:    itemID,
		Mask:      mask,
		MaskHash:  maskHash,
		CreatedBy: createdBy,
	}, nil
}

func generateHash(mask string) string {
	h := sha256.Sum256([]byte(mask))
	return hex.EncodeToString(h[:])[:8]
}

package entity

import (
	"errors"

	"github.com/shopspring/decimal"
)

var (
	ErrInvalidBomID        = errors.New("bom_id must be greater than zero")
	ErrInvalidComponentID  = errors.New("component_id must be greater than zero")
	ErrInvalidQuantity     = errors.New("quantity must be greater than zero")
	ErrInvalidUom          = errors.New("uom must be greater than zero")
	ErrInvalidScrapPercent = errors.New("scrap_percent must be greater than zero")
	ErrInvalidOperationID  = errors.New("operation_id must be greater than zero")
)

func NewBomItems(
	bom_id int64,
	component_id int64,
	quantity decimal.Decimal,
	uom string,
	scrap_percent decimal.Decimal,
	operation_id int64,
	mask_component int64,
) (*BomItems, error) {
	switch {
	case bom_id < 0:
		return nil, ErrInvalidBomID
	case component_id < 0:
		return nil, ErrInvalidComponentID
	case quantity.LessThanOrEqual(decimal.Zero):
		return nil, ErrInvalidQuantity
	case uom == "":
		return nil, ErrInvalidUom
	case scrap_percent.LessThan(decimal.Zero):
		return nil, ErrInvalidScrapPercent
	case operation_id < 0:
		return nil, ErrInvalidOperationID
	}

	return &BomItems{
		BomID:         bom_id,
		ComponentID:   component_id,
		Quantity:      quantity,
		Uom:           uom,
		ScrapPercent:  scrap_percent,
		OperationID:   operation_id,
		MaskComponent: mask_component,
	}, nil
}

package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type BomItems struct {
	ID            int64
	BomID         int64
	ComponentID   int64
	Quantity      decimal.Decimal
	Uom           string
	ScrapPercent  decimal.Decimal
	OperationID   int64
	CreatedAt     time.Time
	MaskComponent int64
}

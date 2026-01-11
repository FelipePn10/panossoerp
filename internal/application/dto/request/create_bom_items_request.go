package request

import (
	"github.com/shopspring/decimal"
)

type CreateBomItemsRequestDTO struct {
	BomID        int64           `json:"bom_id"`
	ComponentID  int64           `json:"component_id"`
	Quantity     decimal.Decimal `json:"quantity"`
	Uom          string          `json:"uom"`
	ScrapPercent string          `json:"scrap_percent"`
	OperationID  int64           `json:"operation_id"`
}

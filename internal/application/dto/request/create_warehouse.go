package request

import (
	"github.com/FelipePn10/panossoerp/internal/domain/enums/types"
	"github.com/google/uuid"
)

type CreateWarehouseRequestDTO struct {
	Code        int    `json:"code"`
	Description string `json:"description"`

	Location types.TypeLocation  `json:"location"`
	Type     types.TypeWarehouse `json:"type"`

	Disposition         bool `json:"disposition"`
	ReservationsAllowed bool `json:"reservations_allowed"`

	CreatedBy uuid.UUID `json:"created_by"`
}

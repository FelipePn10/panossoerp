package entity

import (
	"errors"

	"github.com/FelipePn10/panossoerp/internal/domain/enums/types"
	"github.com/google/uuid"
)

func NewWarehouse(
	code int,
	description string,
	location types.TypeLocation,
	types types.TypeWarehouse,
	disposition bool,
	reservationsAllowed bool,
	created_by uuid.UUID,
) (*Warehouse, error) {
	switch {
	case description == "":
		return nil, errors.ErrUnsupported

	case created_by == uuid.Nil:
		return nil, errors.New("createdby cannot be nil UUID")
	}
	return &Warehouse{
		Code:                code,
		Description:         description,
		Location:            location,
		Type:                types,
		Disposition:         disposition,
		ReservationsAllowed: reservationsAllowed,
		CreatedBy:           created_by,
	}, nil
}

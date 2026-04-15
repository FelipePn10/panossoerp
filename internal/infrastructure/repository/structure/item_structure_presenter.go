package structure

import (
	"github.com/FelipePn10/panossoerp/internal/application/dto/response"
	"github.com/FelipePn10/panossoerp/internal/domain/structure/entity"
)

func ToItemStructureDTO(e *entity.ItemStructure) *response.StructureComponentResponse {
	if e == nil {
		return nil
	}

	return &response.StructureComponentResponse{
		ID:                e.ID,
		ParentItemID:      e.ParentItemID,
		ChildItemID:       e.ChildItemID,
		ParentMask:        e.ParentMask,
		Quantity:          e.Quantity,
		UnitOfMeasurement: e.UnitOfMeasurement,
		LossPercentage:    e.LossPercentage,
		Position:          e.Position,
		Notes:             e.Notes,
		IsActive:          e.IsActive,
		CreatedBy:         e.CreatedBy,
		CreatedAt:         e.CreatedAt,
		UpdatedAt:         e.UpdatedAt,
	}
}

func ToItemStructureListDTO(items []*entity.ItemStructure) []*response.StructureComponentResponse {
	result := make([]*response.StructureComponentResponse, 0, len(items))

	for _, item := range items {
		result = append(result, ToItemStructureDTO(item))
	}

	return result
}

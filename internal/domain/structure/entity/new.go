package entity

import (
	"errors"
	"time"

	"github.com/FelipePn10/panossoerp/internal/domain/enums/types"
	"github.com/google/uuid"
)

func NewItemStructure(
	parentCode, childCode int64,
	parentMask *string,
	quantity float64,
	uom types.TypeUnitOfMeasurementItem,
	health types.Health,
	lossPercentage float64,
	sequence int,
	notes *string,
	isActive bool,
	createdBy uuid.UUID,
) (*ItemStructure, error) {
	if parentCode <= 0 {
		return nil, errors.New("parent_code deve ser positivo")
	}
	if childCode <= 0 {
		return nil, errors.New("child_item_id deve ser positivo")
	}
	if quantity <= 0 {
		return nil, errors.New("quantity deve ser maior que zero")
	}
	if lossPercentage < 0 || lossPercentage > 100 {
		return nil, errors.New("loss_percentage deve estar entre 0 e 100")
	}
	if sequence < 1 {
		sequence = 10
	}
	if parentMask != nil && *parentMask == "" {
		return nil, errors.New("parent_mask não pode ser uma string vazia; use nil para genérico")
	}

	return &ItemStructure{
		ParentCode:        parentCode,
		ChildCode:         childCode,
		ParentMask:        parentMask,
		Quantity:          quantity,
		UnitOfMeasurement: uom,
		Health:            health,
		LossPercentage:    lossPercentage,
		Sequence:          sequence,
		Notes:             notes,
		IsActive:          isActive,
		CreatedBy:         createdBy,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}, nil
}

// IsGeneric retorna true quando o componente se aplica a TODAS as configurações.
func (s *ItemStructure) IsGeneric() bool {
	return s.ParentMask == nil
}

// EffectiveQuantity retorna a quantidade já considerando o percentual de perda.
// Ex.: 10 unidades com 5% de perda → 10.5
func (s *ItemStructure) EffectiveQuantity() float64 {
	return s.Quantity * (1 + s.LossPercentage/100.0)
}

// Deactivate realiza o soft-delete do componente.
func (s *ItemStructure) Deactivate() {
	s.IsActive = false
	s.UpdatedAt = time.Now()
}

func (s *ItemStructure) Update(quantity float64, uom types.TypeUnitOfMeasurementItem, health types.Health, lossPercentage float64, sequence int, notes *string) error {
	if quantity <= 0 {
		return errors.New("quantity deve ser maior que zero")
	}
	if lossPercentage < 0 || lossPercentage > 100 {
		return errors.New("loss_percentage deve estar entre 0 e 100")
	}
	if sequence < 1 {
		sequence = 10
	}
	s.Quantity = quantity
	s.UnitOfMeasurement = uom
	s.Health = health
	s.LossPercentage = lossPercentage
	s.Sequence = sequence
	s.Notes = notes
	s.UpdatedAt = time.Now()
	return nil
}

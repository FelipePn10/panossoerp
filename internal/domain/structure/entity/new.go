package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// NewItemStructure cria e valida uma nova instância de ItemStructure.
func NewItemStructure(
	parentItemID, childItemID int64,
	parentCode, childCode string,
	parentMask *string,
	quantity float64,
	uom string,
	lossPercentage float64,
	position int,
	notes *string,
	createdBy uuid.UUID,
) (*ItemStructure, error) {
	if parentItemID <= 0 {
		return nil, errors.New("parent_item_id deve ser positivo")
	}
	if parentCode == "" {
		return nil, errors.New("parent_code deve ser positivo")
	}
	if childCode <= "" {
		return nil, errors.New("child_item_id deve ser positivo")
	}
	if childItemID <= 0 {
		return nil, errors.New("child_item_id deve ser positivo")
	}
	if parentItemID == childItemID {
		return nil, errors.New("um item não pode ser componente de si mesmo")
	}
	if quantity <= 0 {
		return nil, errors.New("quantity deve ser maior que zero")
	}
	if uom == "" {
		return nil, errors.New("unit_of_measurement é obrigatório")
	}
	if lossPercentage < 0 || lossPercentage > 100 {
		return nil, errors.New("loss_percentage deve estar entre 0 e 100")
	}
	if position < 1 {
		position = 1
	}
	if parentMask != nil && *parentMask == "" {
		return nil, errors.New("parent_mask não pode ser uma string vazia; use nil para genérico")
	}

	return &ItemStructure{
		ParentItemID:      parentItemID,
		ParentCode:        parentCode,
		ChildItemID:       childItemID,
		ChildCode:         childCode,
		ParentMask:        parentMask,
		Quantity:          quantity,
		UnitOfMeasurement: uom,
		LossPercentage:    lossPercentage,
		Position:          position,
		Notes:             notes,
		IsActive:          true,
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

func (s *ItemStructure) Update(quantity float64, uom string, lossPercentage float64, position int, notes *string) error {
	if quantity <= 0 {
		return errors.New("quantity deve ser maior que zero")
	}
	if uom == "" {
		return errors.New("unit_of_measurement é obrigatório")
	}
	if lossPercentage < 0 || lossPercentage > 100 {
		return errors.New("loss_percentage deve estar entre 0 e 100")
	}
	if position < 1 {
		position = 1
	}
	s.Quantity = quantity
	s.UnitOfMeasurement = uom
	s.LossPercentage = lossPercentage
	s.Position = position
	s.Notes = notes
	s.UpdatedAt = time.Now()
	return nil
}

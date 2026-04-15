package request

import "github.com/google/uuid"

// CreateStructureComponentDTO representa a entrada para criar um componente
// de estrutura (BOM).
//
// Regras:
//   - ParentMask nil  → componente genérico (aplica-se a todas as configurações)
//   - ParentMask != nil → componente específico para aquela configuração
type CreateStructureComponentDTO struct {
	ParentItemID      int64     `json:"parent_item_id"`
	ParentCode        string    `json:"parent_code"`
	ChildItemID       int64     `json:"child_item_id"`
	ChildCode         string    `json:"child_code"`
	ParentMask        *string   `json:"parent_mask,omitempty"` // nil = genérico
	Quantity          float64   `json:"quantity"`
	UnitOfMeasurement string    `json:"unit_of_measurement"`
	LossPercentage    float64   `json:"loss_percentage"`
	Position          int       `json:"position"`
	Notes             *string   `json:"notes,omitempty"`
	CreatedBy         uuid.UUID `json:"created_by"`
}

type UpdateStructureComponentDTO struct {
	ID                int64   `json:"id"`
	Quantity          float64 `json:"quantity"`
	UnitOfMeasurement string  `json:"unit_of_measurement"`
	LossPercentage    float64 `json:"loss_percentage"`
	Position          int     `json:"position"`
	Notes             *string `json:"notes,omitempty"`
}

// GetStructureTreeDTO representa a entrada para buscar a árvore BOM genérica
// de um item (sem resolução de máscara).
type GetStructureTreeDTO struct {
	RootItemID int64 `json:"root_item_id"`
}

// ResolveStructureForMaskDTO representa a entrada para resolver a árvore BOM
// completa de um item para uma configuração específica (máscara).
//
// A máscara é propagada automaticamente do pai para os filhos com base
// nas perguntas compartilhadas.
type ResolveStructureForMaskDTO struct {
	RootItemID    int64  `json:"root_item_id"`
	RootMaskValue string `json:"root_mask_value"` // ex: "100#100#50"
}

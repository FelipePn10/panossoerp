package response

import (
	"time"

	"github.com/google/uuid"
)

// StructureComponentResponse representa um único componente de estrutura
// (relação pai → filho) na resposta da API.
type StructureComponentResponse struct {
	ID                int64     `json:"id"`
	ParentItemID      int64     `json:"parent_item_id"`
	ParentCode        string    `json:"parent_code"`
	ChildItemID       int64     `json:"child_item_id"`
	ChildCode         string    `json:"child_code"`
	ChildDescription  string    `json:"child_description"`
	ParentMask        *string   `json:"parent_mask,omitempty"`
	IsGeneric         bool      `json:"is_generic"`
	Quantity          float64   `json:"quantity"`
	EffectiveQuantity float64   `json:"effective_quantity"` // quantity + loss
	UnitOfMeasurement string    `json:"unit_of_measurement"`
	LossPercentage    float64   `json:"loss_percentage"`
	Position          int       `json:"position"`
	Notes             *string   `json:"notes,omitempty"`
	IsActive          bool      `json:"is_active"`
	CreatedBy         uuid.UUID `json:"created_by"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// StructureTreeNodeResponse representa um nó na árvore BOM serializada,
// incluindo os filhos recursivamente.
type StructureTreeNodeResponse struct {
	// Dados do componente (relação pai→filho)
	Component StructureComponentResponse `json:"component"`
	// Máscara calculada para este nó (apenas em modo resolved)
	ResolvedMask *string `json:"resolved_mask,omitempty"`
	// Profundidade na árvore (0 = primeiro nível abaixo da raiz)
	Level int `json:"level"`
	// Filhos resolvidos recursivamente
	Children []*StructureTreeNodeResponse `json:"children"`
}

// StructureTreeResponse é a resposta completa da árvore BOM de um item.
type StructureTreeResponse struct {
	RootItemID  int64                        `json:"root_item_id"`
	RootCode    string                       `json:"root_code"`
	RootDesc    string                       `json:"root_description"`
	RootMask    *string                      `json:"root_mask,omitempty"`
	Components  []*StructureTreeNodeResponse `json:"components"` // nós de 1º nível
	TotalLevels int                          `json:"total_levels"`
	TotalNodes  int                          `json:"total_nodes"`
}

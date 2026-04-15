package entity

import (
	"time"

	"github.com/google/uuid"
)

// ItemStructure representa um componente dentro de uma estrutura de produto (BOM).
//
// Regras de negócio:
//   - ParentMask == nil  → componente genérico: aplica-se a TODAS as configurações
//   - ParentMask != nil  → componente específico: aplica-se APENAS à máscara informada
//   - Um item não pode ser componente de si mesmo
//   - A adição de um componente não pode criar um ciclo na árvore
type ItemStructure struct {
	ID                int64
	ParentItemID      int64
	ParentCode        string
	ChildItemID       int64
	ChildCode         string
	ParentMask        *string // nil = genérico
	Quantity          float64
	UnitOfMeasurement string
	LossPercentage    float64 // 0–100 (%)
	Position          int
	Notes             *string
	IsActive          bool
	CreatedBy         uuid.UUID
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

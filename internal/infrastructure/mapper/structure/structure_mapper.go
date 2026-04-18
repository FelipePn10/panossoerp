package mapper

import (
	"github.com/FelipePn10/panossoerp/internal/application/dto/response"
	"github.com/FelipePn10/panossoerp/internal/domain/structure/valueobject"
)

// domínio → response
// MapNodeToResponse converte um StructureNode do domínio para DTO de resposta.
func MapNodeToResponse(node *valueobject.StructureNode) *response.StructureTreeNodeResponse {
	if node == nil || node.Component == nil {
		return nil
	}

	comp := node.Component
	cr := response.StructureComponentResponse{
		ID:                comp.ID,
		ParentItemCode:    node.ItemCode,
		ChildItemCode:     node.ItemCode,
		ChildDescription:  node.ItemDesc,
		ParentMask:        comp.ParentMask,
		IsGeneric:         comp.IsGeneric(),
		Quantity:          comp.Quantity,
		EffectiveQuantity: comp.EffectiveQuantity(),
		UnitOfMeasurement: comp.UnitOfMeasurement,
		Health:            comp.Health,
		LossPercentage:    comp.LossPercentage,
		Position:          comp.Sequence,
		Notes:             comp.Notes,
		IsActive:          comp.IsActive,
		CreatedBy:         comp.CreatedBy,
		CreatedAt:         comp.CreatedAt,
		UpdatedAt:         comp.UpdatedAt,
	}

	resp := &response.StructureTreeNodeResponse{
		Component:    cr,
		ResolvedMask: node.ResolvedMask,
		Level:        node.Level,
		Children:     make([]*response.StructureTreeNodeResponse, 0, len(node.Children)),
	}

	for _, child := range node.Children {
		if mapped := MapNodeToResponse(child); mapped != nil {
			resp.Children = append(resp.Children, mapped)
		}
	}

	return resp
}

// CountNodes conta recursivamente o total de nós em uma lista de respostas.
func CountNodes(nodes []*response.StructureTreeNodeResponse) int {
	total := 0
	for _, n := range nodes {
		total++
		total += CountNodes(n.Children)
	}
	return total
}

// MaxLevel retorna a profundidade máxima de uma lista de nós.
func MaxLevel(nodes []*response.StructureTreeNodeResponse) int {
	m := 0
	for _, n := range nodes {
		if n.Level > m {
			m = n.Level
		}
		if childMax := MaxLevel(n.Children); childMax > m {
			m = childMax
		}
	}
	return m
}

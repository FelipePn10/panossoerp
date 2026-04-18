package valueobject

import (
	"fmt"
	"strings"

	"github.com/FelipePn10/panossoerp/internal/domain/structure/entity"
)

// StructureNode representa um nó da árvore BOM já resolvida.
// Carrega o componente, informações do item filho, a máscara calculada
// para esse nível e os nós filhos recursivos.
type StructureNode struct {
	// Component é o relacionamento pai→filho (quantity, mask, etc.)
	Component *entity.ItemStructure

	// ItemCode e ItemDesc identificam o item filho de forma legível.
	ItemCode int64
	ItemDesc string

	// Level é a profundidade na árvore (1 = primeiro nível abaixo da raiz).
	Level int

	// ResolvedMask é a máscara calculada para ESTE nó com base nos option_ids
	// propagados do pai. nil quando o item não tem perguntas configuradas
	// ou quando estamos em modo genérico (sem máscara).
	ResolvedMask *string

	// Children são os filhos deste nó, já resolvidos recursivamente.
	Children []*StructureNode
}

// NewStructureNode cria um novo nó da árvore BOM.
func NewStructureNode(
	component *entity.ItemStructure,
	code int64,
	desc string,
	level int,
	resolvedMask *string,
) *StructureNode {
	return &StructureNode{
		Component:    component,
		ItemCode:     code,
		ItemDesc:     desc,
		Level:        level,
		ResolvedMask: resolvedMask,
		Children:     make([]*StructureNode, 0),
	}
}

// AddChild adiciona um filho ao nó atual.
func (n *StructureNode) AddChild(child *StructureNode) {
	n.Children = append(n.Children, child)
}

// IsLeaf retorna true quando o nó não possui filhos.
func (n *StructureNode) IsLeaf() bool {
	return len(n.Children) == 0
}

// HasResolvedMask retorna true quando a máscara deste nó foi calculada.
func (n *StructureNode) HasResolvedMask() bool {
	return n.ResolvedMask != nil && *n.ResolvedMask != ""
}

// MaskAnswer representa a resposta de uma pergunta dentro de uma máscara.
// OptionID referencia a opção selecionada na tabela de opções de pergunta.
// A máscara é formada pela concatenação dos OptionIDs em ordem de posição,
// separados por "#" (ex.: "3#7#2" onde 3, 7 e 2 são option_ids).
type MaskAnswer struct {
	QuestionID int64
	Position   int32
	OptionID   int64
}

// ItemQuestion representa uma pergunta associada a um item.
type ItemQuestion struct {
	QuestionID int64
	Position   int32
}

// PropagateMask calcula a máscara do filho com base nas respostas do pai.
//
// A máscara é uma string de option_ids separados por "#", em ordem de posição
// das perguntas do filho. Ex.: filho tem Q1 (pos 1) e Q2 (pos 2); pai tem
// Q1=optionID 3 e Q2=optionID 7 → máscara do filho = "3#7".
//
// Algoritmo:
//  1. Monta um mapa questionID → optionID a partir das respostas do pai.
//  2. Para cada pergunta do filho (em ordem de posição), busca o optionID
//     no mapa do pai.
//  3. Concatena os IDs encontrados com "#", formando a máscara do filho.
//
// Retorna nil quando:
//   - O filho não tem perguntas configuradas.
//   - Alguma pergunta do filho não existe nas respostas do pai
//     (máscara incompleta → fallback para genérico).
func PropagateMask(parentAnswers []MaskAnswer, childQuestions []ItemQuestion) *string {
	if len(childQuestions) == 0 {
		return nil
	}

	// Mapa: questionID → optionID da resposta do pai
	answerMap := make(map[int64]int64, len(parentAnswers))
	for _, a := range parentAnswers {
		answerMap[a.QuestionID] = a.OptionID
	}

	parts := make([]string, 0, len(childQuestions))
	for _, q := range childQuestions {
		optionID, ok := answerMap[q.QuestionID]
		if !ok {
			// Pergunta do filho não está coberta pelas respostas do pai:
			// não é possível construir a máscara completa do filho.
			// Retorna nil para que o use case aplique o fallback genérico.
			return nil
		}
		parts = append(parts, fmt.Sprintf("%d", optionID))
	}

	if len(parts) == 0 {
		return nil
	}

	mask := strings.Join(parts, "#")
	return &mask
}

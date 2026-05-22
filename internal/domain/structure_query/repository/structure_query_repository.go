package repository

import (
	"context"
	"time"

	maskservice "github.com/FelipePn10/panossoerp/internal/domain/generate_mask_for_item/mask/service"
	maskvo "github.com/FelipePn10/panossoerp/internal/domain/generate_mask_for_item/valueobject"
	str "github.com/FelipePn10/panossoerp/internal/domain/structure/entity"
	"github.com/google/uuid"
)

type StructureQueryRepository interface {
	// Estrutura — mask="" retorna apenas filhos universais (parent_mask IS NULL).
	// mask="1.94M#1.94M" retorna universais + específicos para aquela máscara.
	GetDirectChildrenForMask(ctx context.Context, parentCode int64, mask string) ([]*str.ItemStructure, error)

	// Máscara — o SQL deve fazer JOIN e retornar o option_value.
	GetMaskAnswersByItemAndValue(ctx context.Context, itemCode int64, mask string) ([]maskvo.MaskAnswer, error)
	GetItemQuestions(ctx context.Context, itemCode int64) ([]maskservice.ItemQuestion, error)

	// Cria automaticamente uma máscara propagada; as respostas já chegam com
	// questionID + optionID + position, sem precisar de reverse-lookup.
	CreateMaskForItem(ctx context.Context, itemCode int64, mask string, answers []maskservice.ChildMaskAnswerInput, createdBy uuid.UUID) error

	// Consulta VENG0401 — retorna filhos com campos de data, fórmula e dados do item filho.
	// mask="" retorna apenas filhos genéricos; effectivenessDate nil desativa filtro de data.
	ConsultChildren(ctx context.Context, parentCode int64, mask string, effectivenessDate *time.Time) ([]*str.ConsultRow, error)

	// GetMaskAnswersWithNames retorna question_name → option_value para avaliação de fórmulas.
	GetMaskAnswersWithNames(ctx context.Context, itemCode int64, mask string) (map[string]float64, error)

	// GetLatestMaskForItem retorna a máscara mais recente do item ou "" se não houver.
	GetLatestMaskForItem(ctx context.Context, itemCode int64) (string, error)

	// GetWhereUsed retorna todos os produtos que contêm o componente itemCode (implosão).
	// levels=0 = todos os níveis; levels=N = até N níveis acima.
	GetWhereUsed(ctx context.Context, itemCode int64, levels int) ([]*str.WhereUsedRow, error)
}

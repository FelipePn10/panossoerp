package entity

import (
	"time"

	"github.com/FelipePn10/panossoerp/internal/domain/enums/types"
	"github.com/google/uuid"
)

type Item struct {
	ID          int64
	Code        string
	Name        string
	Description string
	Complement  *string
	Generic     bool
	Configured  bool
	ItemBase    bool
	Process     bool

	//---- PDM
	GroupID    int32
	ModifierID int32

	Situation types.TypeSituationItem
	Health    types.Health

	CreatedBy uuid.UUID
	CreatedAt time.Time

	// --- Pastas

	// Almoxarifado
	WarehouseID         int32
	UnitOfMeasurement   types.TypeUnitOfMeasurementItem // Qual unidade de medida será armazenada para esse item
	AutomaticLow        bool                            // Faz baixa autom?
	CyclicalCount       bool                            // Contagem cíclica
	CyclicalCountConfig *CyclicalCountConfig
	// Engenharia
	ItemBaseCod int
	GrossWeight int16
	NetWeight   int16
	CubicVolume int16
	Type        types.TypeItem
	TypeStruct  types.TypeStructItem
	OEM         bool // componentes ou produtos que são montados sob a marca de outra empresa e revendidos pela empresa contratante do sistema

	//Status    types.Status
}

type CyclicalCountConfig struct {
	DaysInterval int   // A cada quantos dias contar
	MinimumStock int32 // Estoque mínimo para alerta
}

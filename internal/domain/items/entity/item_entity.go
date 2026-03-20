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
	AverageConsumption  bool // Calcular consumo médio/mês

	// Engenharia
	ItemBaseCod int
	GrossWeight int16
	NetWeight   int16
	CubicVolume int16
	Type        types.TypeItem
	TypeStruct  types.TypeStructItem
	OEM         bool // componentes ou produtos que são montados sob a marca de outra empresa e revendidos pela empresa contratante do sistema

	// Planejamento
	TypeMRP            types.TypeMRPItem
	LLC                int // niveis 9 sendo para matérias primas, 2 há 8 para estruras e conjuntos por ex e 1 para o produto final (Que é vendido)
	ReorderPoint       bool
	ReorderPointStruct *ReorderPointStruct
	TankID             int16 // Setor onde é feito

	//Status    types.Status
}

type CyclicalCountConfig struct {
	DaysInterval int   // A cada quantos dias contar
	MinimumStock int32 // Estoque mínimo para alerta
}

type ReorderPointStruct struct {
	// PR = (TR x CM / CR) + ES
	TR *int16
	CM *int16
	CR *int
	ES *int16
}

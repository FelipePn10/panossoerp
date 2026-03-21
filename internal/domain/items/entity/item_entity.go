package entity

import (
	"time"

	"github.com/FelipePn10/panossoerp/internal/domain/enums/types"
	"github.com/FelipePn10/panossoerp/internal/domain/machine/entity"
	"github.com/google/uuid"
)

type Item struct {
	ID                   int64
	Code                 string
	DescriptionTechnique string
	Complement           *string
	Generic              bool // item genérico
	Configured           bool // configurado -> pode sofrer variações
	ItemBase             bool // item base que servirá  para outros itens

	//---- PDM
	GroupID    int32  // Grupo de um item, ex: CHAPAS, AÇÕS etc
	ModifierID int32  // Compor a descrição do item, ex: Grupo: CHAPAS Modificador: Chapa Aço Retax
	Attributes string // "nome" para compor, ex: Grupo: CHAPAS Modificador: Chapa Aço Retax Nome: Retax 5MM

	Situation types.TypeSituationItem
	Health    types.Health

	CreatedBy uuid.UUID
	CreatedAt time.Time

	// --- Pastas:

	// Almoxarifado
	WarehouseID                     int32
	UnitOfMeasurement               types.TypeUnitOfMeasurementItem // Qual unidade de medida será armazenada para esse item
	AutomaticLow                    bool                            // Faz baixa autom?
	CyclicalCount                   bool                            // Contagem cíclica
	CyclicalCountConfig             *CyclicalCountConfig
	MinimumStock                    int32 // Estoque mínimo para alerta de compra
	AverageMonthlyConsumption       bool  // Calcular consumo médio/mês
	AverageMonthlyConsumptionManual *int  // Consumo médio mensal inserido manualmente, apenas se AverageMonthlyConsumption for false

	// Engenharia
	ItemBaseCod *int // Somente se ItemBase for false
	GrossWeight int16
	NetWeight   int16

	Measurements bool // Se true:
	Length       *int
	DepthOrWidth *int
	Height       *int
	CubicVolume  *int16 // É gerado automaticamente com base no comprimento, profundidade e altura do produto

	Type       types.TypeItem
	TypeStruct types.TypeStructItem
	OEM        bool // componentes ou produtos que são montados sob a marca de outra empresa e revendidos pela empresa contratante do sistema

	// Planejamento
	// Para o MRP calcular e gerar ordem de máteria prima, o nivél deve ser LLC 9 e ser GHOST
	TypeMRP            types.TypeMRPItem
	LLC                int // niveis 1 para o produto final, 2 há 8 para estruras e conjuntos e 9 sendo para matérias primas
	ReorderPoint       bool
	ReorderPointStruct *ReorderPointStruct
	TankID             int // Setor onde é feito
	Ghost              bool

	// Planejadores
	EmployeeID    *int32 // Funcionário
	OccupyMachine *bool
	Machines      []entity.MachineUsage

	// Suprimentos
	TypeOfUse types.TypeOfUseItem

	//Status    types.Status
}

type CyclicalCountConfig struct {
	DaysInterval int // A cada quantos dias contar
}

type ReorderPointStruct struct {
	// PR = (TR x CM / CR) + ES
	TR *int16
	CM *int16
	CR *int
	ES *int16
}

package request

type CreateItemDTO struct {
	Code       string  `json:"code"`
	Complement *string `json:"complement,omitempty"`

	Nature int `json:"nature"`

	PDM PDMDTO `json:"pdm"`

	Situation int `json:"situation"`
	Health    int `json:"health"`

	Warehouse   WarehouseDTO   `json:"warehouse"`
	Engineering EngineeringDTO `json:"engineering"`
	Planning    PlanningDTO    `json:"planning"`
	Planners    PlannersDTO    `json:"planners"`
	Supplies    SuppliesDTO    `json:"supplies"`
}

type PDMDTO struct {
	GroupID              int32    `json:"group_id"`
	ModifierID           int32    `json:"modifier_id"`
	Attributes           []string `json:"attributes"` // simplificado
	DescriptionTechnique string   `json:"description_technique"`
}

type WarehouseDTO struct {
	WarehouseID                     int                     `json:"warehouse_id"`
	UnitOfMeasurement               int                     `json:"unit_of_measurement"`
	AutomaticLow                    bool                    `json:"automatic_low"`
	CyclicalCountConfig             *CyclicalCountConfigDTO `json:"cyclical_count_config,omitempty"`
	MinimumStock                    int32                   `json:"minimum_stock"`
	AverageMonthlyConsumptionManual *int                    `json:"average_monthly_consumption_manual,omitempty"`
}

type EngineeringDTO struct {
	ItemBaseCod *int `json:"item_base_cod,omitempty"`

	Weight     float64        `json:"weight"`
	Dimensions *DimensionsDTO `json:"dimensions,omitempty"`

	Type       int  `json:"type"`
	TypeStruct int  `json:"type_struct"`
	OEM        bool `json:"oem"`
}

type PlanningDTO struct {
	TypeMRP      int              `json:"type_mrp"`
	LLC          int              `json:"llc"`
	ReorderPoint *ReorderPointDTO `json:"reorder_point,omitempty"`
	TankID       *int             `json:"tank_id,omitempty"`
	Ghost        bool             `json:"ghost"`
}

type PlannersDTO struct {
	EmployeeID *int32             `json:"employee_id,omitempty"`
	MachinesID *[]MachineUsageDTO `json:"machines,omitempty"`
}

type SuppliesDTO struct {
	TypeOfUse int `json:"type_of_use"`
}

type DimensionsDTO struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Depth  float64 `json:"depth"`
}

type ReorderPointDTO struct {
	Quantity int32 `json:"quantity"`
}

type CyclicalCountConfigDTO struct {
	Frequency int `json:"frequency"`
}

type MachineUsageDTO struct {
	MachineID int `json:"machine_id"`
	Days      int `json:"days"`
}

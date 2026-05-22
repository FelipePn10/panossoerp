package request

type UpsertWorkCenterCostDTO struct {
	WorkCenterID int64   `json:"work_center_id"`
	CostPerHour  float64 `json:"cost_per_hour"`
	Currency     string  `json:"currency"`
	UpdatedBy    string  `json:"updated_by"`
}

type UpsertItemPurchaseCostDTO struct {
	ItemCode  int64   `json:"item_code"`
	UnitCost  float64 `json:"unit_cost"`
	Currency  string  `json:"currency"`
	UpdatedBy string  `json:"updated_by"`
}

type CostRollupDTO struct {
	ItemCode     int64  `json:"item_code"`
	Mask         string `json:"mask"`
	CalculatedBy string `json:"calculated_by"`
}

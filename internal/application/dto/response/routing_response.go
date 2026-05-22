package response

import "time"

type OperationResponse struct {
	ID                  int64     `json:"id"`
	Code                int64     `json:"code"`
	Name                string    `json:"name"`
	Description         *string   `json:"description,omitempty"`
	Origin              string    `json:"origin"`
	Situation           string    `json:"situation"`
	DefaultWorkCenterID *int64    `json:"default_work_center_id,omitempty"`
	StandardTime        float64   `json:"standard_time"`
	SetupTime           float64   `json:"setup_time"`
	IsActive            bool      `json:"is_active"`
	CreatedAt           time.Time `json:"created_at"`
}

type ManufacturingRouteResponse struct {
	ID          int64     `json:"id"`
	Code        int64     `json:"code"`
	ItemCode    int64     `json:"item_code"`
	Mask        *string   `json:"mask,omitempty"`
	Alternative int16     `json:"alternative"`
	Description *string   `json:"description,omitempty"`
	Situation   string    `json:"situation"`
	IsStandard  bool      `json:"is_standard"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
}

type RouteOperationResponse struct {
	ID               int64    `json:"id"`
	RouteID          int64    `json:"route_id"`
	Sequence         int16    `json:"sequence"`
	OperationID      int64    `json:"operation_id"`
	OperationName    string   `json:"operation_name"`
	WorkCenterID     *int64   `json:"work_center_id,omitempty"`
	WorkCenterName   string   `json:"work_center_name,omitempty"`
	StandardTime     *float64 `json:"standard_time,omitempty"`
	SetupTime        *float64 `json:"setup_time,omitempty"`
	EffectiveStdTime float64  `json:"effective_std_time"`
	EffectiveSetup   float64  `json:"effective_setup"`
	Situation        string   `json:"situation"`
	Notes            *string  `json:"notes,omitempty"`
}

type NetworkEdgeResponse struct {
	ID            int64   `json:"id"`
	PredecessorID int64   `json:"predecessor_id"`
	SuccessorID   int64   `json:"successor_id"`
	OverlapPct    float64 `json:"overlap_pct"`
}

type RouteDetailResponse struct {
	Route      ManufacturingRouteResponse `json:"route"`
	Operations []RouteOperationResponse   `json:"operations"`
	Network    []NetworkEdgeResponse      `json:"network"`
}

type RouteLeadTimeResponse struct {
	RouteID      int64   `json:"route_id"`
	TotalHours   float64 `json:"total_hours"`
	CriticalPath []int64 `json:"critical_path"` // route_operation IDs
}

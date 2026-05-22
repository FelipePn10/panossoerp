package response

import "time"

type CRPSummaryResponse struct {
	PlanCode      int64 `json:"plan_code"`
	TotalEntries  int   `json:"total_entries"`
	OverloadCount int   `json:"overload_count"`
}

type CRPEntryResponse struct {
	ID             int64     `json:"id"`
	PlanCode       int64     `json:"plan_code"`
	WorkCenterID   int64     `json:"work_center_id"`
	ReqDate        time.Time `json:"req_date"`
	RequiredHours  float64   `json:"required_hours"`
	AvailableHours float64   `json:"available_hours"`
	LoadPct        float64   `json:"load_pct"`
	IsOverloaded   bool      `json:"is_overloaded"`
}

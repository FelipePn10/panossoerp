package entity

import "time"

type CapacityRequirement struct {
	ID             int64
	PlanCode       int64
	WorkCenterID   int64
	ReqDate        time.Time
	RequiredHours  float64
	AvailableHours float64
	LoadPct        float64
	CreatedAt      time.Time
}

type WorkCenterLoad struct {
	WorkCenterID   int64
	ReqDate        time.Time
	RequiredHours  float64
	AvailableHours float64
	LoadPct        float64
	IsOverloaded   bool
}

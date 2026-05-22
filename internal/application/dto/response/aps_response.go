package response

import "time"

type APSSummaryResponse struct {
	ScheduledOperations int `json:"scheduled_operations"`
	OrdersProcessed     int `json:"orders_processed"`
}

type GanttTaskResponse struct {
	SequenceID        int64     `json:"sequence_id"`
	ProductionOrderID int64     `json:"production_order_id"`
	WorkCenterID      int64     `json:"work_center_id"`
	SequencePosition  int       `json:"sequence_position"`
	ScheduledStart    time.Time `json:"scheduled_start"`
	ScheduledEnd      time.Time `json:"scheduled_end"`
	Status            string    `json:"status"`
	DurationHours     float64   `json:"duration_hours"`
}

package request

import "time"

type SequenceOrdersDTO struct {
	StartFrom time.Time `json:"start_from"`
}

type GanttByWorkCenterDTO struct {
	WorkCenterID int64     `json:"work_center_id"`
	From         time.Time `json:"from"`
	To           time.Time `json:"to"`
}

package entity

import "time"

type MachineUsage struct {
	ID        int64
	ItemID    int64
	MachineID int
	UsageTime int
	CreatedAt time.Time
}

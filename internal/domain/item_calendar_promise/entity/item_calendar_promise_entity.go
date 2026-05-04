package entity

import "time"

type ItemCalendarPromise struct {
	ID          int64
	ItemCode    int64
	Mask        string
	Year        int
	Month       int
	Day         int
	IsWorkday   bool
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

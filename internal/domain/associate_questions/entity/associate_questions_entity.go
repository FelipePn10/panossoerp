package entity

import "time"

type AssociateQuestion struct {
	ItemID     int64
	QuestionID int64
	Position   int
	CreatedAt  time.Time
}

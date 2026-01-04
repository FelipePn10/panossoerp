package entity

import "time"

type AssociateQuestion struct {
	ProductID  int64
	QuestionID int64
	Position   int
	CreatedAt  time.Time
}

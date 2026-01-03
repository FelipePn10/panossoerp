package entity

import "time"

type ProductQuestion struct {
	ProductID  int64
	QuestionID int64
	Position   int
	CreatedAt  time.Time
}

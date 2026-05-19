package entity

import "time"

type AssociateQuestion struct {
	ItemCode   int64
	QuestionID int64
	Position   int
	CreatedAt  time.Time
}

type AssociateQuestionDetail struct {
	ItemCode     int64
	QuestionID   int64
	QuestionName string
	Position     int
	CreatedAt    time.Time
}

type ItemQuestionRow struct {
	ItemCode         int64
	ItemBusinessCode int64
	QuestionID       int64
	QuestionName     string
	Position         int
	CreatedAt        time.Time
}

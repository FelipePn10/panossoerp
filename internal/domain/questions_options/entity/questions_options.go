package entity

import "github.com/google/uuid"

type QuestionsOptions struct {
	ID         int64
	QuestionId int64
	CreatedBy  uuid.UUID
	Value      string
}

package entity

import "github.com/google/uuid"

type QuestionsOptions struct {
	QuestionId int64
	CreatedBy  uuid.UUID
	Value      string
}

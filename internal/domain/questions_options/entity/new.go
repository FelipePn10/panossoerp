package entity

import (
	"errors"

	"github.com/google/uuid"
)

func NewQuestionsOptions(
	value string,
	questionId int64,
	createdBy uuid.UUID,
) (*QuestionsOptions, error) {
	switch {
	case value == "":
		return nil, errors.ErrUnsupported
	case createdBy == uuid.Nil:
		return nil, errors.New("createdby cannot be nil UUID")
	}

	return &QuestionsOptions{
		Value:      value,
		QuestionId: questionId,
		CreatedBy:  createdBy,
	}, nil
}

func ValidateQuestionOptionDeletion(id int64) error {
	if id <= 0 {
		return errors.New("product id must be greater than zero")
	}
	return nil
}

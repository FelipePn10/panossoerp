package entity

import (
	"errors"

	"github.com/google/uuid"
)

func NewQuestion(
	name string,
	createdBy uuid.UUID,
) (*Question, error) {
	switch {
	case name == "":
		return nil, errors.ErrUnsupported
	case createdBy == uuid.Nil:
		return nil, errors.New("createdby cannot be nil UUID")
	}

	return &Question{
		Name:      name,
		CreatedBy: createdBy,
	}, nil
}

func ValidateQuestionDeletion(id int64) error {
	if id <= 0 {
		return errors.New("product id must be greater than zero")
	}
	return nil
}

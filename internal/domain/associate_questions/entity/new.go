package entity

import (
	"errors"
	"time"
)

var (
	ErrInvalidPosition = errors.New("position must be greater than zero")
)

func New(
	productID int64,
	questionId int64,
	position int,
) (*ProductQuestion, error) {
	if position <= 0 {
		return nil, ErrInvalidPosition
	}

	return &ProductQuestion{
		ProductID:  productID,
		QuestionID: questionId,
		Position:   position,
		CreatedAt:  time.Now(),
	}, nil
}

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
) (*AssociateQuestion, error) {
	if position <= 0 {
		return nil, ErrInvalidPosition
	}

	return &AssociateQuestion{
		ProductID:  productID,
		QuestionID: questionId,
		Position:   position,
		CreatedAt:  time.Now(),
	}, nil
}

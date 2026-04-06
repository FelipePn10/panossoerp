package entity

import (
	"errors"
	"time"
)

var (
	ErrInvalidPosition = errors.New("position must be greater than zero")
)

func New(
	item int64,
	questionId int64,
	position int,
) (*AssociateQuestion, error) {
	if position <= 0 {
		return nil, ErrInvalidPosition
	}

	return &AssociateQuestion{
		ItemID:     item,
		QuestionID: questionId,
		Position:   position,
		CreatedAt:  time.Now(),
	}, nil
}

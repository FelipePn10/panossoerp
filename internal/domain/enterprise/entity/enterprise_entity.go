package entity

import "github.com/google/uuid"

type Enterprise struct {
	ID        int
	Code      int
	Name      string
	CreatedBy uuid.UUID
}

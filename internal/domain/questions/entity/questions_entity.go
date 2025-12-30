package entity

import (
	"github.com/google/uuid"
)

type Question struct {
	Name      string
	CreatedBy uuid.UUID
}

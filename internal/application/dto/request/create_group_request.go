package request

import "github.com/google/uuid"

type CreateGroupDTO struct {
	Code         int       `json:"code"`
	Description  string    `json:"description"`
	EnterpriseID int       `json:"enterprise_id"`
	CreatedBy    uuid.UUID `json:"created_by"`
}

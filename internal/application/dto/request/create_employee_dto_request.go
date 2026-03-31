package request

type CreateEmployeeDTO struct {
	EnterpriseID int    `json:"enterprise_id"`
	Code         int    `json:"code"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}

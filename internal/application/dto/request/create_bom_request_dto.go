package request

import "time"

type CreateBomUseCaseRequestDTO struct {
	ProductId int64     `json:"product_id"`
	BomType   string    `json:"bom_type"`
	Version   int       `json:"version"`
	Status    string    `json:"status"`
	ValidFrom time.Time `json:"valid_from"`
}

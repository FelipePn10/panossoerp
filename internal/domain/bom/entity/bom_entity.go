package entity

import "time"

type Bom struct {
	ID        int64
	ProductId int64
	BomType   string
	Version   int
	Status    string
	ValidFrom time.Time
	CreatedAt time.Time
}

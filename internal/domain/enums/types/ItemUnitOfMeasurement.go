package types

import (
	"encoding/json"
)

type TypeUnitOfMeasurementItem int

const (
	MM TypeUnitOfMeasurementItem = iota
	CM
	M
	IN
	KG
	M2
	M3
	UN
	MICROMETRO
	TONELADA
)

func (t TypeUnitOfMeasurementItem) String() string {
	switch t {
	case MM:
		return "MM"
	case CM:
		return "CM"
	case M:
		return "M"
	case IN:
		return "IN"
	case KG:
		return "KG"
	case M2:
		return "M2"
	case M3:
		return "M3"
	case UN:
		return "UN"
	case MICROMETRO:
		return "MICROMETRO"
	case TONELADA:
		return "TONELADA"

	default:
		return "Desconhecido"
	}
}

func (t TypeUnitOfMeasurementItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

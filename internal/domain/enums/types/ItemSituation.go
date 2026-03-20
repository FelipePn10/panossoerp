package types

import "encoding/json"

type TypeSituationItem int

const (
	LINHA TypeSituationItem = iota
	PROMOCAO
)

func (t TypeSituationItem) String() string {
	switch t {
	case LINHA:
		return "LINHA"
	case PROMOCAO:
		return "PROMOCAO"

	default:
		return "Desconhecido"
	}
}

func (t TypeSituationItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

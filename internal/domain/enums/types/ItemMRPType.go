package types

import "encoding/json"

type TypeMRPItem int

const (
	NORMAL_MRP TypeMRPItem = iota
	PROJETO
)

func (t TypeMRPItem) String() string {
	switch t {
	case NORMAL_MRP:
		return "NORMAL_MRP"
	case PROJETO:
		return "PROJETO"

	default:
		return "Desconhecido"
	}
}

func (t TypeMRPItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

package types

import "encoding/json"

type TypeWarehouse int

const (
	LINHA_DE_PRODUCAO TypeWarehouse = iota
	NORMAL
)

func (t TypeWarehouse) String() string {
	switch t {
	case LINHA_DE_PRODUCAO:
		return "LINHA DE PRODUÇÃO"
	case NORMAL:
		return "NORMAL"

	default:
		return "Desconhecido"
	}
}

func (t TypeWarehouse) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

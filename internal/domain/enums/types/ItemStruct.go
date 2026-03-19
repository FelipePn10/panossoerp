package types

import "encoding/json"

type TypeStructItem int

const (
	INDUSTRIAL TypeStructItem = iota // Itens do qual o MRP gera ordem e controla estoque
	COMERCIAL                        // Item pronto para a venda
)

func (t TypeStructItem) String() string {
	switch t {
	case INDUSTRIAL:
		return "INDUSTRIAL"
	case COMERCIAL:
		return "COMERCIAL"

	default:
		return "Desconhecido"
	}
}

func (t TypeStructItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

package types

import "encoding/json"

type TypeLocation int

const (
	INTERNO TypeLocation = iota
	EXTERNO
	ASSISTENCIA
	REJEICAO
	INSPECAO
	RESERVA
	TRANSITO
	ESPECIAL
)

func (t TypeLocation) String() string {
	switch t {
	case INTERNO:
		return "INTERNO"
	case EXTERNO:
		return "EXTERNO"
	case ASSISTENCIA:
		return "ASSISTÊNCIA"
	case REJEICAO:
		return "REJEIÇÃO"
	case INSPECAO:
		return "INSPEÇÃO"
	case RESERVA:
		return "RESERVA"
	case TRANSITO:
		return "TRÂNSITO"
	case ESPECIAL:
		return "ESPECIAL"
	default:
		return "NENHUM"
	}
}

func (t TypeLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

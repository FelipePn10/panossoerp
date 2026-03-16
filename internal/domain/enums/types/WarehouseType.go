package types

type TypeWarehouse int

const (
	MATERIA_PRIMA TypeWarehouse = iota
	CHAPAS_METALICAS
	PERFIS_METALICOS
	TUBOS_METALICOS
	FERRAGENS_FIXADORES
	SOLDAGEM
	PINTURA_ACABAMENTO
	FERRAMENTAS
	MANUTENCAO
	PRODUTO_SEMIACABADO
	PRODUTO_ACABADO
	EXPEDICAO
	ALMOXARIFADO_GERAL
)

func (t TypeWarehouse) String() string {
	switch t {
	case MATERIA_PRIMA:
		return "Matéria-prima"
	case CHAPAS_METALICAS:
		return "Chapas metálicas"
	case PERFIS_METALICOS:
		return "Perfis metálicos"
	case TUBOS_METALICOS:
		return "Tubos metálicos"
	case FERRAGENS_FIXADORES:
		return "Ferragens e fixadores"
	case SOLDAGEM:
		return "Soldagem"
	case PINTURA_ACABAMENTO:
		return "Pintura e acabamento"
	case FERRAMENTAS:
		return "Ferramentas"
	case MANUTENCAO:
		return "Manutenção"
	case PRODUTO_SEMIACABADO:
		return "Produto semiacabado"
	case PRODUTO_ACABADO:
		return "Produto acabado"
	case EXPEDICAO:
		return "Expedição"
	case ALMOXARIFADO_GERAL:
		return "Almoxarifado geral"
	default:
		return "Desconhecido"
	}
}

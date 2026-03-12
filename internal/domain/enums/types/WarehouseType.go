package types

type TypeWarehouse int

const (
	WOOD = iota
	FOAMS
	ACCESSORIES
	PAINTS
	FEET
	GENERAL
)

func (s TypeWarehouse) String() string {
	switch s {
	case WOOD:
		return "WOOD"
	case FOAMS:
		return "FOAMS"
	case ACCESSORIES:
		return "ACCESSORIES"
	case PAINTS:
		return "PAINTS"
	case FEET:
		return "FEET"
	case GENERAL:
		return "GENERAL"
	default:
		return "UNKNOWN"
	}
}

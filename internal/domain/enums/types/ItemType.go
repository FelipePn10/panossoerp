package types

type Type int

const (
	SERVICE = iota
	RAW_MATERIAL
	COMPONENT
	SEMI_FINISHED
	FINISHED_PRODUCT
	CONSUMABLE
)

func (s Type) String() string {
	switch s {
	case SERVICE:
		return "SERVICE"
	case RAW_MATERIAL:
		return "RAW_MATERIAL"
	case COMPONENT:
		return "COMPONENT"
	case SEMI_FINISHED:
		return "SEMI_FINISHED"
	case FINISHED_PRODUCT:
		return "FINISHED_PRODUCT"
	case CONSUMABLE:
		return "CONSUMABLE"
	default:
		return "UNKNOWN"
	}
}

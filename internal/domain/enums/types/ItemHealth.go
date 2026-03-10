package types

type Health int

const (
	ACTIVE = iota
	INACTIVE
	GHOST
)

func (s Health) String() string {
	switch s {
	case ACTIVE:
		return "ACTIVE"
	case INACTIVE:
		return "INACTIVE"
	case GHOST:
		return "GHOST"
	default:
		return "UNKNOWN"
	}
}

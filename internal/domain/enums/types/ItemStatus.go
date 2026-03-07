package types

type Status int

const (
	MANUFACTURED = iota
	PURCHASED
)

func (s Status) String() string {
	switch s {
	case MANUFACTURED:
		return "MANUFACTURED"
	case PURCHASED:
		return "PURCHASED"
	default:
		return "UNKNOWN"
	}
}

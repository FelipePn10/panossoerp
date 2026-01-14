package valueobject

import (
	"errors"
	"fmt"
	"math/rand"
)

type ComponentCode struct {
	value string
}

func NewComponentCode(groupCode string) (ComponentCode, error) {
	if len(groupCode) < 2 {
		return ComponentCode{}, errors.New("group code must have at least 2 characters")
	}

	group := groupCode[:2]
	random := rand.Intn(10000)

	code := fmt.Sprintf("%s%04d", group, random)
	return ComponentCode{value: code}, nil
}

func (c ComponentCode) String() string {
	return c.value
}

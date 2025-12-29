package valueobject

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type ProductCode struct {
	value string
}

func NewProductCode(groupCode string, createdAt time.Time) (ProductCode, error) {
	if len(groupCode) < 2 {
		return ProductCode{}, errors.New("group code must have at least 2 characters")
	}

	group := groupCode[:2]
	date := createdAt.Format("0201")
	random := rand.Intn(100)

	code := fmt.Sprintf("%s%s%02d", group, date, random)

	return ProductCode{value: code}, nil
}

func (p ProductCode) String() string {
	return p.value
}

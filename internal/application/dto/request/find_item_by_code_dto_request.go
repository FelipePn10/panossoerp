package request

import "github.com/FelipePn10/panossoerp/internal/domain/items/valueobject"

type FindItemByCodeDTO struct {
	Code valueobject.ItemCode `json:"code"`
}

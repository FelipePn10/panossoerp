package valueobject

import "errors"

type MaskForProduct struct {
	mask string
}

func NewMaskForProduct(productCode string, questionsAnswered string) (MaskForProduct, error) {
	if len(productCode) < 10 || productCode == "" {
		return MaskForProduct{}, errors.New("invalid product code")
	}

}

package valueobject

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type MaskAnswer struct {
	questionID int64
	optionID   int64
	position   int
}

type ProductMask struct {
	productCode string
	createdBy   uuid.UUID
	answers     []MaskAnswer
	mask        string
}

func NewMaskAnswer(questionID, optionID int64, position int) (MaskAnswer, error) {
	if questionID <= 0 {
		return MaskAnswer{}, errors.New("invalid question id")
	}
	if optionID <= 0 {
		return MaskAnswer{}, errors.New("invalid option id")
	}
	if position <= 0 {
		return MaskAnswer{}, errors.New("invalid position")
	}

	return MaskAnswer{
		questionID: questionID,
		optionID:   optionID,
		position:   position,
	}, nil
}

func NewProductMask(productCode string, createdBy uuid.UUID, answers []MaskAnswer) (ProductMask, error) {
	if productCode == "" {
		return ProductMask{}, errors.New("invalid product code")
	}
	if len(answers) == 0 {
		return ProductMask{}, errors.New("mask must have at least one answer")
	}

	mask := generateMask(answers)

	return ProductMask{
		productCode: productCode,
		createdBy:   createdBy,
		answers:     answers,
		mask:        mask,
	}, nil
}

func generateMask(answers []MaskAnswer) string {
	sort.Slice(answers, func(i, j int) bool {
		return answers[i].position < answers[j].position
	})

	values := make([]string, 0, len(answers))
	for _, a := range answers {
		values = append(values, strconv.FormatInt(a.optionID, 10))
	}

	return strings.Join(values, "#")
}

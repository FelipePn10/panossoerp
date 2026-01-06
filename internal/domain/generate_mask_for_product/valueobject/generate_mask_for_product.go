package valueobject

import (
	"crypto/sha256"
	"encoding/hex"
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
	hash        string
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

	h := sha256.Sum256([]byte(mask))
	hash := hex.EncodeToString(h[:])[:8]

	return ProductMask{
		productCode: productCode,
		createdBy:   createdBy,
		answers:     answers,
		mask:        mask,
		hash:        hash,
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

func (pm ProductMask) Value() string {
	return pm.mask
}

func (pm ProductMask) Hash() string {
	return pm.hash
}

func (ma MaskAnswer) QuestionID() int64 {
	return ma.questionID
}

func (ma MaskAnswer) OptionID() int64 {
	return ma.optionID
}

func (ma MaskAnswer) Position() int {
	return ma.position
}

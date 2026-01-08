package errorsuc

import "errors"

var ErrQuestionNotFound = errors.New("question not found")
var ErrInvalidQuestionName = errors.New("invalid question name")
var ErrInvalidProductNameAndCodeNotFound = errors.New("name and/or code not found.")
var ErrProductNotFound = errors.New("product not found")

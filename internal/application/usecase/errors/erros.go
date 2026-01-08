package errorsuc

import "errors"

var ErrQuestionNotFound = errors.New("question not found")
var ErrInvalidQuestionName = errors.New("invalid question name")

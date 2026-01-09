package errorsuc

import "errors"

var ErrQuestionNotFound = errors.New("question not found")
var ErrInvalidQuestionName = errors.New("invalid question name")
var ErrInvalidProductNameAndCodeNotFound = errors.New("name and/or code not found.")
var ErrProductNotFound = errors.New("product not found")
var ErrCreateBom = errors.New("there was an error while trying to register the good.")
var ErrCreateBomNotFound = errors.New("Unable to complete the order. Please check if you are sending the correct data using the correct method and URL.")

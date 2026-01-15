package errorsuc

import "errors"

var (
	ErrQuestionNotFound                  = errors.New("question not found")
	ErrQuestionOptionAlreadyExists       = errors.New("question option already exists")
	ErrQuestionAlreadyExists             = errors.New("question option already exists")
	ErrInvalidQuestionName               = errors.New("invalid question name")
	ErrInvalidProductNameAndCodeNotFound = errors.New("name and/or code not found.")
	ErrProductNotFound                   = errors.New("product not found")
	ErrProductAlreadyExists              = errors.New("product already exists")
	ErrCreateBom                         = errors.New("there was an error while trying to register the good.")
	ErrCreateBomNotFound                 = errors.New("Unable to complete the order. Please check if you are sending the correct data using the correct method and URL.")
	ErrCreateBomItem                     = errors.New("there was an error while trying to register the good.")
	ErrBomItemAlreadyExists              = errors.New("product already exists")
	ErrCreateBomItemNotFound             = errors.New("Unable to complete the order. Please check if you are sending the correct data using the correct method and URL.")
	ErrComponentAlreadyExists            = errors.New("component already exists")
	ErrWarehouseNotFound                 = errors.New("warehouse not found")
	ErrUnauthorized                      = errors.New("user not authorized")
	ErrInvalidSearchParams               = errors.New("params invalid")
)

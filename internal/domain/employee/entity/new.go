package entity

import "errors"

var (
	ErrInvalidCode = errors.New("invalid code")
)

func NewEmployee(
	enterprise_id int,
	code int,
	name string,
	description string,
) (*Employee, error) {
	if code <= 0 {
		return nil, ErrInvalidCode
	}

	employee := &Employee{
		EnterpriseID: enterprise_id,
		Name:         name,
		Description:  description,
		Code:         code,
	}

	return employee, nil
}

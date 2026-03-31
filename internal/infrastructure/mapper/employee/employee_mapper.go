package mapper

import (
	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/domain/employee/entity"
)

func ToEmployeeEntity(d request.CreateEmployeeDTO) (*entity.Employee, error) {
	return entity.NewEmployee(
		d.EnterpriseID,
		d.Code,
		d.Description,
		d.Name,
	)
}

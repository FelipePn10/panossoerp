package mapper

import (
	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/domain/enterprise/entity"
)

func ToEnterpriseEntity(d request.CreateEnterpriseDTO) (*entity.Enterprise, error) {
	return entity.NewEnterprise(
		d.Code,
		d.Name,
		d.CreatedBy,
	)
}

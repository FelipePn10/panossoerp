package mapper

import (
	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/domain/group/entity"
)

func ToGroupEntity(d request.CreateGroupDTO) (*entity.Group, error) {
	return entity.NewGroup(
		d.Code,
		d.Description,
		d.EnterpriseID,
		d.CreatedBy,
	)
}

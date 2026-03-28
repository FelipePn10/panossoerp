package mapper

import (
	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	"github.com/FelipePn10/panossoerp/internal/domain/modifier/entity"
)

func ToModifierEntity(d request.CreateModifierDTO) (*entity.Modifier, error) {
	return entity.NewModifier(
		d.Description,
		d.CreatedBy,
	)
}

package structure_uc

import (
	"context"
	"fmt"

	"github.com/FelipePn10/panossoerp/internal/application/dto/response"
	sqrepo "github.com/FelipePn10/panossoerp/internal/domain/structure_query/repository"
)

type WhereUsedUseCase struct {
	repo sqrepo.StructureQueryRepository
}

func NewWhereUsedUseCase(repo sqrepo.StructureQueryRepository) *WhereUsedUseCase {
	return &WhereUsedUseCase{repo: repo}
}

func (uc *WhereUsedUseCase) Execute(ctx context.Context, itemCode int64, levels int) (*response.WhereUsedResponse, error) {
	if itemCode <= 0 {
		return nil, fmt.Errorf("item_code must be positive")
	}
	rows, err := uc.repo.GetWhereUsed(ctx, itemCode, levels)
	if err != nil {
		return nil, err
	}

	out := make([]response.WhereUsedRowResponse, 0, len(rows))
	for _, r := range rows {
		row := response.WhereUsedRowResponse{
			Level:             r.Level,
			ParentCode:        r.ParentCode,
			ParentDescription: r.ParentDescription,
			ChildCode:         r.ChildCode,
			Quantity:          r.Quantity,
			LossPercentage:    r.LossPercentage,
			Sequence:          r.Sequence,
			ParentMask:        r.ParentMask,
		}
		out = append(out, row)
	}
	return &response.WhereUsedResponse{ItemCode: itemCode, Rows: out}, nil
}

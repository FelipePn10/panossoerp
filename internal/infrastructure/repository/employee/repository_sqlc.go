package employee

import (
	"context"
	"fmt"

	"github.com/FelipePn10/panossoerp/internal/domain/employee/entity"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

func (r *repositoryEmployeeSQLC) Create(
	ctx context.Context,
	employee *entity.Employee,
) (*entity.Employee, error) {
	params := sqlc.CreateEmployeeParams{}

	dbEmployee, err := r.q.CreateEmployee(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("create employee: %w", err)
	}

	return &entity.Employee{}, nil
}

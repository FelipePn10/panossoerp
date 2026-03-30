package repository

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/domain/employee/entity"
)

type EmployeeRepository interface {
	Create(ctx context.Context, employee *entity.Employee) (*entity.Employee, error)
}

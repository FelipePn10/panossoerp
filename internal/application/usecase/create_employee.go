package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/ports"
	errorsuc "github.com/FelipePn10/panossoerp/internal/application/usecase/errors"
	"github.com/FelipePn10/panossoerp/internal/domain/employee/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/employee/repository"
)

type CreateEmployeeUseCase struct {
	repo repository.EmployeeRepository
	auth ports.AuthService
}

func (uc *CreateEmployeeUseCase) Execute(
	ctx context.Context,
	employee *entity.Employee,
) (*entity.Employee, error) {
	if !uc.auth.CanCreateEmployee(ctx) {
		return nil, errorsuc.ErrUnauthorized
	}

	created, err := uc.repo.Create(ctx, employee)
	if err != nil {
		return nil, err
	}

	return created, nil
}

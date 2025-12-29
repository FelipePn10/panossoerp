package usecase

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/dto/request"
	user "github.com/FelipePn10/panossoerp/internal/domain/user/entity"
	"github.com/FelipePn10/panossoerp/internal/domain/user/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserUseCase struct {
	repo repository.UserRepository
}

func (uc *RegisterUserUseCase) Execute(
	ctx context.Context,
	dto request.RegisterUserDTO,
) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user, err := user.NewUser(
		uuid.New(),
		dto.Name,
		dto.Email,
		string(hash),
	)
	if err != nil {
		return err
	}

	return uc.repo.Create(ctx, user)
}

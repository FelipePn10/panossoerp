package usecase

import (
	"github.com/FelipePn10/panossoerp/internal/domain/product/repository"
	user "github.com/FelipePn10/panossoerp/internal/domain/user/repository"
)

func NewCreateProductUseCase(
	repo repository.ProductRepository,
) *CreateProductUseCase {
	return &CreateProductUseCase{
		repo: repo,
	}
}

func NewDeleteProductUseCase(
	repo repository.ProductRepository,
) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		repo: repo,
	}
}

// func NewSearchByIDProductUseCase(
// 	repo repository.ProductRepository,
// ) *SearchByIDProductUseCase {
// 	return &SearchByIDProductUseCase{
// 		repo: repo,
// 	}
// }

func NewLoginUserUseCase(repo user.UserRepository) *LoginUserUseCase {
	return &LoginUserUseCase{repo: repo}
}

func NewRegisterUserUseCase(
	repo user.UserRepository,
) *RegisterUserUseCase {
	return &RegisterUserUseCase{repo: repo}
}

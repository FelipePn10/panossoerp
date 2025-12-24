package handler

import "github.com/FelipePn10/panossoerp/internal/application/usecase"

func NewCreateProductHandler(createProductUC *usecase.CreateProductUseCase) *ProductHandler {
	return &ProductHandler{
		createProductUC: createProductUC,
	}
}

func NewUserHandler(
	registerUC *usecase.RegisterUserUseCase,
	loginUC *usecase.LoginUserUseCase,
	jwtSecret string,
) *UserHandler {
	return &UserHandler{
		registerUC: registerUC,
		loginUC:    loginUC,
		jwtSecret:  jwtSecret,
	}
}

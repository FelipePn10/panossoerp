package handler

import "github.com/FelipePn10/panossoerp/internal/application/usecase"

type ProductHandler struct {
	createProductUC *usecase.CreateProductUseCase
}

type UserHandler struct {
	registerUC *usecase.RegisterUserUseCase
	loginUC    *usecase.LoginUserUseCase
	jwtSecret  string
}

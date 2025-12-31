package handler

import "github.com/FelipePn10/panossoerp/internal/application/usecase"

type ProductHandler struct {
	createProductUC *usecase.CreateProductUseCase
	deleteProductUC *usecase.DeleteProductUseCase
}

type UserHandler struct {
	registerUC *usecase.RegisterUserUseCase
	loginUC    *usecase.LoginUserUseCase
	jwtSecret  string
}

type QuestionHandler struct {
	createQuestionUC *usecase.CreateQuestion
	deleteQuestionUC *usecase.DeleteQuestionUseCase
}

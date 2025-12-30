package handler

import "github.com/FelipePn10/panossoerp/internal/application/usecase"

func NewCreateProductHandler(
	createProductUC *usecase.CreateProductUseCase,
) *ProductHandler {
	return &ProductHandler{
		createProductUC: createProductUC,
	}
}

func NewDeleteProductHandler(
	deleteProductUC *usecase.DeleteProductUseCase,
) *ProductHandler {
	return &ProductHandler{
		deleteProductUC: deleteProductUC,
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

func NewQuestionHandler(
	createQuestionUC *usecase.CreateQuestion,
) *QuestionHandler {
	return &QuestionHandler{
		createQuestionUC: createQuestionUC,
	}
}

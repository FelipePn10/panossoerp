package handler

import (
	"github.com/FelipePn10/panossoerp/internal/application/usecase"
)

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

func NewFindProductByNameAndCodeHandler(
	findProductByNameAndCodeUC *usecase.FindProductByNameAndCode,
) *ProductHandler {
	return &ProductHandler{
		findProductByNameAndCodeUC: findProductByNameAndCodeUC,
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

func NewDeleteQuestionHandler(
	deleteQuestionUC *usecase.DeleteQuestionUseCase,
) *QuestionHandler {
	return &QuestionHandler{
		deleteQuestionUC: deleteQuestionUC,
	}
}

func NewCreateQuestionOptionHandler(
	createQuestionOptionUC *usecase.CreateQuestionOptionUseCase,
) *QuestionOptionHandler {
	return &QuestionOptionHandler{
		createQuestionOptionUC: createQuestionOptionUC,
	}
}

func NewDeleteQuestionOptionHandler(
	deleteQuestionOptionUC *usecase.DeleteQuestionOptionUseCase,
) *QuestionOptionHandler {
	return &QuestionOptionHandler{
		deleteQuestionOptionUC: deleteQuestionOptionUC,
	}
}

func NewAssociateByQuestionProductHandler(
	associateByQuestionProductUC *usecase.AssociateByQuestionProductUseCase,
) *AssociateByQuestionProductHandler {
	return &AssociateByQuestionProductHandler{
		associateByQuestionProductUC: associateByQuestionProductUC,
	}
}

func NewGeneratMaskProductHandler(
	generateMaskProductUC *usecase.GenerateMaskForProductUseCase,
) *GenerateMaskHandler {
	return &GenerateMaskHandler{
		generateMask: generateMaskProductUC,
	}
}

package handler

import "github.com/FelipePn10/panossoerp/internal/application/usecase"

type ProductHandler struct {
	createProductUC            *usecase.CreateProductUseCase
	deleteProductUC            *usecase.DeleteProductUseCase
	findProductByNameAndCodeUC *usecase.FindProductByNameAndCode
}

type UserHandler struct {
	registerUC *usecase.RegisterUserUseCase
	loginUC    *usecase.LoginUserUseCase
	jwtSecret  string
}

type QuestionHandler struct {
	createQuestionUC     *usecase.CreateQuestion
	deleteQuestionUC     *usecase.DeleteQuestionUseCase
	findQuestionByNameUC *usecase.FindQuestionByName
}

type QuestionOptionHandler struct {
	createQuestionOptionUC *usecase.CreateQuestionOptionUseCase
	deleteQuestionOptionUC *usecase.DeleteQuestionOptionUseCase
}

type AssociateByQuestionProductHandler struct {
	associateByQuestionProductUC *usecase.AssociateByQuestionProductUseCase
}

type GenerateMaskHandler struct {
	generateMask *usecase.GenerateMaskForProductUseCase
}

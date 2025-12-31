package usecase

import (
	prdt "github.com/FelipePn10/panossoerp/internal/domain/product/repository"
	qst "github.com/FelipePn10/panossoerp/internal/domain/questions/repository"
	user "github.com/FelipePn10/panossoerp/internal/domain/user/repository"
)

func NewCreateProductUseCase(
	repo prdt.ProductRepository,
) *CreateProductUseCase {
	return &CreateProductUseCase{
		repo: repo,
	}
}

func NewDeleteProductUseCase(
	repo prdt.ProductRepository,
) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		repo: repo,
	}
}

func NewFindProductByNameAndCode(
	repo prdt.ProductRepository,
) *FindProductByNameAndCode {
	return &FindProductByNameAndCode{
		repo: repo,
	}
}

func NewDeleteQuestionUseCase(
	repo qst.QuestionsRepository,
) *DeleteQuestionUseCase {
	return &DeleteQuestionUseCase{
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

func NewLoginUserUseCase(
	repo user.UserRepository,
) *LoginUserUseCase {
	return &LoginUserUseCase{repo: repo}
}

func NewRegisterUserUseCase(
	repo user.UserRepository,
) *RegisterUserUseCase {
	return &RegisterUserUseCase{repo: repo}
}

func NewQuestionUserUseCase(
	repo qst.QuestionsRepository,
) *CreateQuestion {
	return &CreateQuestion{repo: repo}
}

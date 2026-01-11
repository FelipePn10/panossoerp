package usecase

import (
	ast "github.com/FelipePn10/panossoerp/internal/domain/associate_questions/repository"
	bom "github.com/FelipePn10/panossoerp/internal/domain/bom/repository"
	bomitem "github.com/FelipePn10/panossoerp/internal/domain/bom_items/repository"
	mask "github.com/FelipePn10/panossoerp/internal/domain/generate_mask_for_product/repository"
	prdt "github.com/FelipePn10/panossoerp/internal/domain/product/repository"
	qst "github.com/FelipePn10/panossoerp/internal/domain/questions/repository"
	qstops "github.com/FelipePn10/panossoerp/internal/domain/questions_options/repository"
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

func NewFindQuestionByName(
	repo qst.QuestionsRepository,
) *FindQuestionByName {
	return &FindQuestionByName{
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

func NewCreateQuestionOptionUseCase(
	repo qstops.QuestionsOptionsRepository,
) *CreateQuestionOptionUseCase {
	return &CreateQuestionOptionUseCase{
		repo: repo,
	}
}
func NewDeleteQuestionOptionUseCase(
	repo qstops.QuestionsOptionsRepository,
) *DeleteQuestionOptionUseCase {
	return &DeleteQuestionOptionUseCase{
		repo: repo,
	}
}

func NewAssociateByQuestionProductUseCase(
	repo ast.AssociateQuestionsRepository,
) *AssociateByQuestionProductUseCase {
	return &AssociateByQuestionProductUseCase{
		repo: repo,
	}
}

func NewGenerateMaskProductUseCase(
	repo mask.GenerateMaskForProductRepository,
) *GenerateMaskForProductUseCase {
	return &GenerateMaskForProductUseCase{
		repo: repo,
	}
}
func NewCreateBomUseCase(
	repo bom.BomRepository,
) *CreateBomUseCase {
	return &CreateBomUseCase{
		repo: repo,
	}
}

func NewCreatBomItemUseCase(
	repo bomitem.BomItemsRepository,
) *CreateBomItemUseCase {
	return &CreateBomItemUseCase{
		repo: repo,
	}
}

package usecase

import (
	"github.com/FelipePn10/panossoerp/internal/application/ports"
	ast "github.com/FelipePn10/panossoerp/internal/domain/associate_questions/repository"
	bom "github.com/FelipePn10/panossoerp/internal/domain/bom/repository"
	bomitem "github.com/FelipePn10/panossoerp/internal/domain/bom_items/repository"
	component "github.com/FelipePn10/panossoerp/internal/domain/component/repository"
	mask "github.com/FelipePn10/panossoerp/internal/domain/generate_mask_for_product/repository"
	prdt "github.com/FelipePn10/panossoerp/internal/domain/product/repository"
	qst "github.com/FelipePn10/panossoerp/internal/domain/questions/repository"
	qstops "github.com/FelipePn10/panossoerp/internal/domain/questions_options/repository"
	user "github.com/FelipePn10/panossoerp/internal/domain/user/repository"
)

func NewCreateProductUseCase(
	repo prdt.ProductRepository,
	auth ports.AuthService,
) *CreateProductUseCase {
	return &CreateProductUseCase{
		repo: repo,
		auth: auth,
	}
}

func NewDeleteProductUseCase(
	repo prdt.ProductRepository,
	auth ports.AuthService,
) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		repo: repo,
		auth: auth,
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
	auth ports.AuthService,
) *CreateQuestion {
	return &CreateQuestion{
		repo: repo,
		auth: auth,
	}
}

func NewCreateQuestionOptionUseCase(
	repo qstops.QuestionsOptionsRepository,
	auth ports.AuthService,
) *CreateQuestionOptionUseCase {
	return &CreateQuestionOptionUseCase{
		repo: repo,
		auth: auth,
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
	auth ports.AuthService,
) *AssociateByQuestionProductUseCase {
	return &AssociateByQuestionProductUseCase{
		repo: repo,
		auth: auth,
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
	auth ports.AuthService,
) *CreateBomUseCase {
	return &CreateBomUseCase{
		repo: repo,
		auth: auth,
	}
}

func NewCreatBomItemUseCase(
	repo bomitem.BomItemsRepository,
	auth ports.AuthService,
) *CreateBomItemUseCase {
	return &CreateBomItemUseCase{
		repo: repo,
		auth: auth,
	}
}

func NewCreateComponentUseCase(
	repo component.ComponentRepository,
	auth ports.AuthService,
) *CreateComponentUseCase {
	return &CreateComponentUseCase{
		repo: repo,
		auth: auth,
	}
}

package handler

import (
	"github.com/FelipePn10/panossoerp/internal/application/usecase"
	"github.com/FelipePn10/panossoerp/internal/interfaces/http/handler/security"
)

type ProductHandler struct {
	*security.BaseHandler
	createProductUC            *usecase.CreateProductUseCase
	deleteProductUC            *usecase.DeleteProductUseCase
	findProductByNameAndCodeUC *usecase.FindProductByNameAndCode
}

type ItemHandler struct {
	*security.BaseHandler
	createItemUC *usecase.CreateItemUseCase
}

type UserHandler struct {
	*security.BaseHandler
	registerUC *usecase.RegisterUserUseCase
	loginUC    *usecase.LoginUserUseCase
	jwtSecret  string
}

type QuestionHandler struct {
	*security.BaseHandler
	createQuestionUC     *usecase.CreateQuestion
	deleteQuestionUC     *usecase.DeleteQuestionUseCase
	findQuestionByNameUC *usecase.FindQuestionByName
}

type QuestionOptionHandler struct {
	*security.BaseHandler
	createQuestionOptionUC *usecase.CreateQuestionOptionUseCase
	deleteQuestionOptionUC *usecase.DeleteQuestionOptionUseCase
}

type AssociateByQuestionProductHandler struct {
	*security.BaseHandler
	associateByQuestionProductUC *usecase.AssociateByQuestionProductUseCase
}

type GenerateMaskHandler struct {
	*security.BaseHandler
	generateMask *usecase.GenerateMaskForProductUseCase
}

type BomHandler struct {
	*security.BaseHandler
	createBomUC *usecase.CreateBomUseCase
}

type BomItemHandler struct {
	*security.BaseHandler
	createBomItemUC *usecase.CreateBomItemUseCase
}

type WarehouseHandler struct {
	*security.BaseHandler
	createWarehouseUC *usecase.CreateWarehouseUseCase
}

type GroupHandler struct {
	*security.BaseHandler
	createGroupUC *usecase.CreateGroupUseCase
}

type EnterpriseHandler struct {
	*security.BaseHandler
	createEnterpriseUC *usecase.CreateEnterpriseUseCase
}

type ModifierHandler struct {
	*security.BaseHandler
	createModifierUC *usecase.CreateModifierUseCase
}

type EmployeeHandler struct {
	*security.BaseHandler
	createEmployeeUC *usecase.CreateEmployeeUseCase
}

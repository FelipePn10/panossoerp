package handler

import "github.com/FelipePn10/panossoerp/internal/application/usecase"

func NewCreateProductHandler(createProductUC *usecase.CreateProductUseCase) *Handler {
	return &Handler{
		createProductUC: createProductUC,
	}
}

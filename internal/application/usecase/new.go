package usecase

import "github.com/FelipePn10/panossoerp/internal/domain/product/repository"

func NewCreateProductUseCase(
	repo repository.ProductRepository,
) *CreateProductUseCase {
	return &CreateProductUseCase{
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

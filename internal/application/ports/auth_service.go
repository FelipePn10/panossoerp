package ports

import (
	"context"
)

type AuthService interface {
	CanCreateComponent(ctx context.Context) bool
	CanCreateProduct(ctx context.Context) bool
	CanCreateBom(ctx context.Context) bool
	CanCreateBomItems(ctx context.Context) bool
	CanAssociateByQuestionProduct(ctx context.Context) bool
	CanCreateQuestion(ctx context.Context) bool
	CanCreateQuestionOption(ctx context.Context) bool
}

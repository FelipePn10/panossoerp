package auth

import (
	"context"
	"strings"

	"github.com/FelipePn10/panossoerp/internal/application/security"
	contextkey "github.com/FelipePn10/panossoerp/internal/interfaces/http/context"
)

type AuthService struct{}

func (a *AuthService) hasWriteRole(ctx context.Context) bool {
	user, ok := ctx.Value(contextkey.UserKey).(*security.AuthUser)
	if !ok {
		return false
	}

	role := strings.ToUpper(strings.TrimSpace(user.Role))
	return role == "ADMIN" || role == "USER"
}

func (a *AuthService) CanCreateComponent(ctx context.Context) bool {
	return a.hasWriteRole(ctx)
}

func (a *AuthService) CanCreateItem(ctx context.Context) bool {
	return a.hasWriteRole(ctx)
}

func (a *AuthService) CanCreateProduct(ctx context.Context) bool {
	return a.hasWriteRole(ctx)
}

func (a *AuthService) CanCreateBom(ctx context.Context) bool {
	return a.hasWriteRole(ctx)
}

func (a *AuthService) CanCreateBomItems(ctx context.Context) bool {
	return a.hasWriteRole(ctx)
}

func (a *AuthService) CanAssociateByQuestionProduct(ctx context.Context) bool {
	return a.hasWriteRole(ctx)
}

func (a *AuthService) CanCreateQuestion(ctx context.Context) bool {
	return a.hasWriteRole(ctx)
}

func (a *AuthService) CanCreateQuestionOption(ctx context.Context) bool {
	return a.hasWriteRole(ctx)
}

func (a *AuthService) CanDeleteProduct(ctx context.Context) bool {
	return a.hasWriteRole(ctx)
}

func (a *AuthService) CanCreateWarehouse(ctx context.Context) bool {
	return a.hasWriteRole(ctx)
}

func (a *AuthService) CanCreateGroup(ctx context.Context) bool {
	return a.hasWriteRole(ctx)
}

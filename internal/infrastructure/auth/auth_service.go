package auth

import (
	"context"

	"github.com/FelipePn10/panossoerp/internal/application/security"
	contextkey "github.com/FelipePn10/panossoerp/internal/interfaces/http/context"
)

type AuthService struct{}

func (a *AuthService) CanCreateComponent(ctx context.Context) bool {
	user, ok := ctx.Value(contextkey.UserKey).(*security.AuthUser)
	if !ok {
		return false
	}

	return user.Role == "admin" || user.Role == "user"
}

package ports

import (
	"context"
)

type AuthService interface {
	CanCreateComponent(ctx context.Context) bool
}

package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID string `json:"sub"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

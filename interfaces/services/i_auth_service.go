package services

import (
	"context"
	"lawyerinyou-backend/token"
)

type IAuthService interface {
	Logout(ctx context.Context, claims token.Claims, Token string) error
}

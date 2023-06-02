package services

import (
	"context"
	"lawyerinyou-backend/models"
	"lawyerinyou-backend/token"
)

type IAuthService interface {
	Logout(ctx context.Context, claims token.Claims, Token string) error
	Login(ctx context.Context, dataLogin *models.LoginForm) (output interface{}, err error)
}

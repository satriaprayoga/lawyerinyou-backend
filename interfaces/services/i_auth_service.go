package services

import (
	"context"
	"lawyerinyou-backend/models"
	"lawyerinyou-backend/token"
)

type IAuthService interface {
	Logout(ctx context.Context, claims token.Claims, Token string) error
	Login(ctx context.Context, dataLogin *models.LoginForm) (output interface{}, err error)
	ForgotPassword(ctx context.Context, dataForgot *models.ForgotForm) (result string, err error)
	GenOTP(ctx context.Context, dataForgot *models.ForgotForm) (result interface{}, err error)
	Register(ctx context.Context, dataRegister models.RegisterForm) (output interface{}, err error)
}

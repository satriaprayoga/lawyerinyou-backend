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
	ResetPassword(ctx context.Context, dataReset *models.ResetPasswd) (err error)
	Register(ctx context.Context, dataRegister models.RegisterForm) (output interface{}, err error)
	VerifyRegister(ctx context.Context, dataVerify models.VerifyForm) (output interface{}, err error)
	Verify(ctx context.Context, dataVerify models.VerifyForm) (err error)
}

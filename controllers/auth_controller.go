package controllers

import (
	"context"
	"fmt"
	"lawyerinyou-backend/interfaces/services"
	middlewares "lawyerinyou-backend/middleware"
	"lawyerinyou-backend/models"
	"lawyerinyou-backend/pkg/form"
	"lawyerinyou-backend/pkg/logging"
	"lawyerinyou-backend/pkg/responses"
	"lawyerinyou-backend/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authUseCase services.IAuthService
}

func NewAuthController(e *echo.Echo, authUseCase services.IAuthService) {
	cont := &AuthController{authUseCase: authUseCase}

	L := e.Group("/user/auth/logout")
	L.Use(middlewares.JWT)
	L.POST("", cont.Logout)

	r := e.Group("/user/auth")
	r.POST("/login", cont.Login)
	r.POST("/forgot", cont.ForgotPassword)
	r.POST("/change_password", cont.ChangePassword)
	r.POST("/verify", cont.Verify)
	r.POST("/register/verify", cont.VerifyRegister)
	r.POST("/register/gen_otp", cont.GenOTP)
	r.POST("/register", cont.Register)

}

func (a *AuthController) Logout(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		resp = responses.Res{R: e}
	)

	claims, err := form.GetClaims(e)
	if err != nil {
		return resp.Response(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	Token := e.Request().Header.Get("Authorization")
	err = a.authUseCase.Logout(ctx, claims, Token)
	if err != nil {
		return resp.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}
	return resp.Response(http.StatusOK, "Ok", nil)
}

func (u *AuthController) Login(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{}
		resp   = responses.Res{R: e}

		loginForm = models.LoginForm{}
	)
	httpCode, errMsg := form.BindAndValid(e, &loginForm)
	logger.Info(utils.Stringify(loginForm))
	if httpCode != 200 {
		return resp.ResponseError(http.StatusBadRequest, errMsg, nil)
	}
	out, err := u.authUseCase.Login(ctx, &loginForm)
	if err != nil {
		return resp.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}

	return resp.Response(http.StatusOK, "Ok", out)

}
func (u *AuthController) Register(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{}
		resp   = responses.Res{R: e}

		regForm = models.RegisterForm{}
	)
	httpCode, errMsg := form.BindAndValid(e, &regForm)
	logger.Info(utils.Stringify(regForm))
	if httpCode != 200 {
		return resp.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	data, err := u.authUseCase.Register(ctx, regForm)
	if err != nil {
		return resp.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}

	return resp.Response(http.StatusOK, "Ok", data)
}

func (u *AuthController) Verify(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	var (
		logger = logging.Logger{}
		resp   = responses.Res{R: e}

		verifyForm = models.VerifyForm{}
	)

	httpCode, errMsg := form.BindAndValid(e, &verifyForm)
	logger.Info(utils.Stringify(verifyForm))
	if httpCode != 200 {
		return resp.ResponseError(http.StatusBadRequest, errMsg, nil)
	}
	err := u.authUseCase.Verify(ctx, verifyForm)
	if err != nil {
		return resp.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}

	return resp.Response(http.StatusOK, "Ok", nil)

}

func (u *AuthController) VerifyRegister(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	var (
		logger = logging.Logger{}
		resp   = responses.Res{R: e}

		verifyForm = models.VerifyForm{}
	)

	httpCode, errMsg := form.BindAndValid(e, &verifyForm)
	logger.Info(utils.Stringify(verifyForm))
	if httpCode != 200 {
		return resp.ResponseError(http.StatusBadRequest, errMsg, nil)
	}
	data, err := u.authUseCase.VerifyRegister(ctx, verifyForm)
	if err != nil {
		return resp.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}

	return resp.Response(http.StatusOK, "Ok", data)

}

func (u *AuthController) ForgotPassword(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	var (
		logger = logging.Logger{}
		resp   = responses.Res{R: e}

		forgotForm = models.ForgotForm{}
	)

	httpCode, errMsg := form.BindAndValid(e, &forgotForm)
	logger.Info(utils.Stringify(forgotForm))
	if httpCode != 200 {
		return resp.ResponseError(http.StatusBadRequest, errMsg, nil)
	}
	OTP, err := u.authUseCase.ForgotPassword(ctx, &forgotForm)
	if err != nil {
		return resp.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}
	result := map[string]interface{}{
		"otp": OTP,
	}
	return resp.Response(http.StatusOK, "Ok", result)
}

func (u *AuthController) ChangePassword(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	var (
		logger = logging.Logger{}
		resp   = responses.Res{R: e}

		forgotForm = models.ResetPasswd{}
	)

	httpCode, errMsg := form.BindAndValid(e, &forgotForm)
	logger.Info(utils.Stringify(forgotForm))
	if httpCode != 200 {
		return resp.ResponseError(http.StatusBadRequest, errMsg, nil)
	}
	err := u.authUseCase.ResetPassword(ctx, &forgotForm)
	if err != nil {
		return resp.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}
	return resp.Response(http.StatusOK, "Ok", "please login")
}

func (u *AuthController) GenOTP(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	var (
		logger = logging.Logger{}
		resp   = responses.Res{R: e}

		forgotForm = models.ForgotForm{}
	)

	httpCode, errMsg := form.BindAndValid(e, &forgotForm)
	logger.Info(utils.Stringify(forgotForm))
	if httpCode != 200 {
		return resp.ResponseError(http.StatusBadRequest, errMsg, nil)
	}
	OTP, err := u.authUseCase.GenOTP(ctx, &forgotForm)
	if err != nil {
		return resp.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}
	result := map[string]interface{}{
		"otp": OTP,
	}
	return resp.Response(http.StatusOK, "Ok", result)
}

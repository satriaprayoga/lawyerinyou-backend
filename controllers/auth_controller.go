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

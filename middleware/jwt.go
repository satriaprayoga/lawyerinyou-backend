package middlewares

import (
	"lawyerinyou-backend/pkg/redis"
	"lawyerinyou-backend/pkg/responses"
	"lawyerinyou-backend/pkg/settings"
	"lawyerinyou-backend/token"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			code     = http.StatusOK
			msg      = ""
			data     interface{}
			jwtToken = c.Request().Header.Get("Authorization")
		)
		data = map[string]string{
			"token": jwtToken,
		}
		if jwtToken == "" {
			code = http.StatusNetworkAuthenticationRequired
			msg = "Auth Token Required"
		} else {
			existToken := redis.GetSession(jwtToken)
			if existToken == "" {
				code = http.StatusUnauthorized
				msg = "Token Failed"
			}
			claims, err := token.ParseToken(jwtToken)
			if err != nil {
				code = http.StatusUnauthorized
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					msg = "Token Expired"
				default:
					msg = "Token Failed"
				}
			} else {
				var issuer = settings.AppConfigSetting.App.Issuer
				valid := claims.VerifyIssuer(issuer, true)
				if !valid {
					code = http.StatusUnauthorized
					msg = "Issuer is not valid"
				}
				c.Set("claims", claims)
			}
		}
		if code != http.StatusOK {
			resp := responses.ResponseModel{
				Msg:  msg,
				Data: data,
			}
			return c.JSON(code, resp)
		}
		return next(c)
	}
}

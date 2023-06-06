package routes

import (
	"lawyerinyou-backend/controllers"
	"lawyerinyou-backend/pkg/database"
	"lawyerinyou-backend/pkg/settings"
	repoimpl "lawyerinyou-backend/repository"
	"lawyerinyou-backend/usecases"
	"time"

	"github.com/labstack/echo/v4"
)

type AppRoutes struct {
	E *echo.Echo
}

func (e *AppRoutes) InitialRouter() {
	timeoutContext := time.Duration(settings.AppConfigSetting.Server.ReadTimeOut) * time.Second

	userRepo := repoimpl.NewLawUserRepo(database.Conn)
	authUseCase := usecases.NewAuthService(userRepo, timeoutContext)

	controllers.NewAuthController(e.E, authUseCase)
}

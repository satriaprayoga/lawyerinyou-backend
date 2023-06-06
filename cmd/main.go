package main

import (
	"fmt"
	"lawyerinyou-backend/pkg/database"
	"lawyerinyou-backend/pkg/logging"
	"lawyerinyou-backend/pkg/redis"
	"lawyerinyou-backend/pkg/settings"
	"lawyerinyou-backend/routes"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	settings.Setup("./config/config.json")
	database.Setup()
	redis.Setup()
	logging.Setup()
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	R := routes.AppRoutes{E: e}
	R.InitialRouter()

	sPort := fmt.Sprintf(":%d", settings.AppConfigSetting.Server.HTTPPort)

	log.Fatal(e.Start(sPort))
}

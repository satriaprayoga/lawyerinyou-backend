package usecases

import (
	"context"
	"lawyerinyou-backend/models"
	"lawyerinyou-backend/pkg/database"
	"lawyerinyou-backend/pkg/logging"
	"lawyerinyou-backend/pkg/redis"
	"lawyerinyou-backend/pkg/settings"
	"lawyerinyou-backend/pkg/utils"
	repoimpl "lawyerinyou-backend/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	settings.Setup("../config/config.json")
	database.Setup()
	redis.Setup()
	logging.Setup()

	var (
		timeOut      = settings.AppConfigSetting.Server.ReadTimeOut
		ctx          = context.Background()
		registerForm models.RegisterForm
	)

	lawRepo := repoimpl.NewLawUserRepo(database.Conn)
	authService := NewAuthService(lawRepo, time.Second*time.Duration(timeOut))
	registerForm.Account = "satria.prayoga@gmail.com"
	registerForm.BirthOfDate, _ = utils.GetDayOfBirth(1987, 05, 04, "2022-03-12")
	registerForm.Name = "Gilang Satria"
	registerForm.Passwd = "asdqwe123"
	registerForm.ConfirmPasswd = "asdqwe123"
	registerForm.UserType = "user"
	data, err := authService.Register(ctx, registerForm)
	require.NoError(t, err)
	require.NotNil(t, data)

}

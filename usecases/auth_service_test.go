package usecases

import (
	"lawyerinyou-backend/pkg/database"
	"lawyerinyou-backend/pkg/logging"
	"lawyerinyou-backend/pkg/redis"
	"lawyerinyou-backend/pkg/settings"
	"testing"
)

func TestRegister(t *testing.T) {
	settings.Setup("../config/config.json")
	database.Setup()
	redis.Setup()
	logging.Setup()

}

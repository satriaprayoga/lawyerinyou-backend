package email

import (
	"lawyerinyou-backend/pkg/logging"
	"lawyerinyou-backend/pkg/settings"
	"lawyerinyou-backend/pkg/utils"
	"testing"
)

func TestEmailRegister(t *testing.T) {
	settings.Setup("../../config/config.json")
	logging.Setup()

	GenCode := utils.GenerateNumber(4)
	mail := Register{
		Email:      "satria.prayoga@gmail.com",
		Name:       "Gilang Satria",
		PasswordCd: GenCode,
	}
	mail.SendRegister()
}

package email

import (
	"fmt"
	"strings"
)

type Register struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	PasswordCd string `json:"password_cd"`
}

func (R *Register) SendRegister() error {
	subjectEmail := "Informasi OTP"
	fmt.Println(subjectEmail)
	err := SendEmail(R.Email, subjectEmail, getVerifyBody(R))
	if err != nil {
		return err
	}
	return nil
}

func getVerifyBody(R *Register) string {
	registerHTML := SendRegister

	registerHTML = strings.ReplaceAll(registerHTML, `{Name}`, R.Name)
	registerHTML = strings.ReplaceAll(registerHTML, `{Email}`, R.Email)
	registerHTML = strings.ReplaceAll(registerHTML, `{OTP}`, R.PasswordCd)
	return registerHTML
}

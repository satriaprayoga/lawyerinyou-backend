package email

import (
	"fmt"
	"strings"
)

type Forgot struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	OTP   string `json:"otp"`
}

func (F *Forgot) SendForgot() error {
	subjectEmail := "Permintaan Lupa Password"
	fmt.Println(subjectEmail)
	err := SendEmail(F.Email, subjectEmail, getInformasiLoginBodyForgot(F))
	if err != nil {
		return err
	}
	return nil
}

func getInformasiLoginBodyForgot(F *Forgot) string {
	verifyHTML := VerifyCode

	verifyHTML = strings.ReplaceAll(verifyHTML, `{Name}`, F.Name)
	verifyHTML = strings.ReplaceAll(verifyHTML, `{Email}`, F.Email)
	verifyHTML = strings.ReplaceAll(verifyHTML, `{OTP}`, F.OTP)
	return verifyHTML
}

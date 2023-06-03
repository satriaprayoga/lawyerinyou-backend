package email

import "fmt"

func SendEmail(to string, subject string, htmlBody string) {
	fmt.Printf("%s %s %s", to, subject, htmlBody)
}

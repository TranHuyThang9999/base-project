package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	email := "my email@gmail.com"
	pass := "pass key"
	to := "tranhuythang9999@gmail.com"
	subject := "Test Email"
	content := "This is a test email sent from Go."

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", email, to, subject, content)

	auth := smtp.PlainAuth("", email, pass, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, email, []string{to}, []byte(message))
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}

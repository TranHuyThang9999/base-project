package utils

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"rices/common/configs"
	customerrors "rices/core/custom_errors"
)

func SendEmail(toAddress, subject, content string) *customerrors.CustomError {
	cf := configs.Get()
	email := cf.Email
	pass := cf.AppKey
	smtpHost := cf.SmtpHost
	smtpPort := cf.SmtpPort

	if toAddress == "" {
		return customerrors.NewError(nil, http.StatusBadRequest, 1, "recipient email address is empty")
	}
	if subject == "" {
		return customerrors.NewError(nil, http.StatusBadRequest, 2, "email subject is empty")
	}
	if content == "" {
		return customerrors.NewError(nil, http.StatusBadRequest, 3, "email content is empty")
	}

	message := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", email, toAddress, subject, content)

	auth := smtp.PlainAuth("", email, pass, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, email, []string{toAddress}, []byte(message))
	if err != nil {
		log.Println("Error sending email:", err)
		return customerrors.NewError(err, http.StatusInternalServerError, 9, "failed to send email")
	}

	return nil
}

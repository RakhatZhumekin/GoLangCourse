package service

import (
	"net/smtp"
	"os"
	"strconv"
)

type EmailConfirmationService interface {
	SendVerificationCode(rn int, email string) error
}

type emailConfirmationService struct {
}

func NewEmailConfirmationService() EmailConfirmationService {
	return &emailConfirmationService{}
}

func (service *emailConfirmationService) SendVerificationCode(rn int, email string) error {
	from := os.Getenv("FROM_EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	to := []string{email}

	verCode := strconv.Itoa(rn)

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Subject: Email Verification"

	body := "Verification code: " + verCode

	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	return smtp.SendMail(address, auth, from, to, message)
}
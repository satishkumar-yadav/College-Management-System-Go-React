package utils

import (
	"fmt"
	"net/smtp"
)

func SendApprovalEmail(to, name string) error {
	auth := smtp.PlainAuth("", "your-email@gmail.com", "your-app-password", "smtp.gmail.com")

	msg := []byte(fmt.Sprintf("Subject: Account Approved\r\n\r\nDear %s,\n\nYour registration is approved. You can now log in.\n\nRegards,\nXYZ College", name))

	return smtp.SendMail("smtp.gmail.com:587", auth, "your-email@gmail.com", []string{to}, msg)
}

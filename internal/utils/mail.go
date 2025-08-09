package utils

import (
	"fmt"
	"net/smtp"
	"path/filepath"
	"cv-url/config"

	"github.com/jordan-wright/email"
)

func SendEmail(from, login, password, server string, port int, userID, reportFile, zipFile string) error {
	e := email.NewEmail()
	e.From = from
	e.To = []string{config.MailTo}
	e.Subject = fmt.Sprintf("Golang Test – %s", userID)
	e.Text = []byte("Автоматическая отправка отчёта")

	if _, err := e.AttachFile(filepath.Clean(reportFile)); err != nil {
		return err
	}
	if _, err := e.AttachFile(filepath.Clean(zipFile)); err != nil {
		return err
	}

	addr := fmt.Sprintf("%s:%d", server, port)
	auth := smtp.PlainAuth("", login, password, server)

	return e.Send(addr, auth)
}

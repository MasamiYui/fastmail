package service

import (
	"fastmail/internal/config"
	"fmt"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	dialer *gomail.Dialer
	from   string
}

// NewEmailService creates a new EmailService with the given configuration.
func NewEmailService(cfg *config.Config) *EmailService {
	d := gomail.NewDialer(cfg.SMTP.Host, cfg.SMTP.Port, cfg.SMTP.User, cfg.SMTP.Pass)
	return &EmailService{
		dialer: d,
		from:   cfg.SMTP.User,
	}
}

// SendEmail sends an email to the specified recipients.
func (s *EmailService) SendEmail(to []string, subject, body string, attachments []string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	for _, attachment := range attachments {
		m.Attach(attachment)
	}

	if err := s.dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

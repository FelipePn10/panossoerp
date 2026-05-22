package notification

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"strings"
)

type SMTPConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	From     string
}

type EmailService struct {
	cfg SMTPConfig
}

func NewEmailService(cfg SMTPConfig) *EmailService {
	return &EmailService{cfg: cfg}
}

// Enabled reports whether SMTP is configured. When false, Send is a no-op.
func (s *EmailService) Enabled() bool {
	return s.cfg.Host != "" && s.cfg.User != ""
}

// Send delivers a plain-text email to one or more recipients.
// Returns nil when SMTP is not configured (graceful degradation).
func (s *EmailService) Send(to []string, subject, body string) error {
	if !s.Enabled() {
		return nil
	}

	msg := buildMIMEMessage(s.cfg.From, to, subject, body)
	addr := net.JoinHostPort(s.cfg.Host, s.cfg.Port)
	auth := smtp.PlainAuth("", s.cfg.User, s.cfg.Password, s.cfg.Host)

	// Try implicit TLS (port 465). On failure, fall back to STARTTLS via smtp.SendMail.
	tlsCfg := &tls.Config{ServerName: s.cfg.Host}
	conn, err := tls.Dial("tcp", addr, tlsCfg)
	if err != nil {
		return smtp.SendMail(addr, auth, s.cfg.From, to, []byte(msg))
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, s.cfg.Host)
	if err != nil {
		return fmt.Errorf("smtp client: %w", err)
	}
	defer client.Quit()

	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("smtp auth: %w", err)
	}
	if err = client.Mail(s.cfg.From); err != nil {
		return fmt.Errorf("smtp MAIL FROM: %w", err)
	}
	for _, r := range to {
		if err = client.Rcpt(r); err != nil {
			return fmt.Errorf("smtp RCPT TO %s: %w", r, err)
		}
	}
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("smtp DATA: %w", err)
	}
	if _, err = fmt.Fprint(w, msg); err != nil {
		return err
	}
	return w.Close()
}

func buildMIMEMessage(from string, to []string, subject, body string) string {
	var sb strings.Builder
	sb.WriteString("From: " + from + "\r\n")
	sb.WriteString("To: " + strings.Join(to, ", ") + "\r\n")
	sb.WriteString("Subject: " + subject + "\r\n")
	sb.WriteString("MIME-Version: 1.0\r\n")
	sb.WriteString("Content-Type: text/plain; charset=UTF-8\r\n")
	sb.WriteString("\r\n")
	sb.WriteString(body)
	return sb.String()
}

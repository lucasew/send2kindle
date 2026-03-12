package main

import (
	"net/mail"
	"net/smtp"
	"strings"

	"github.com/scorredoira/email"
)

// Email encapsulates the routing and payload structure of a single message,
// consisting of a destination address and an array of absolute local file paths to attach.
type Email struct {
	To    string
	Files []string
}

// EmailAuthenticationData holds the required configuration for the outbound SMTP connection.
// 'Server' usually includes the domain and port (e.g. smtp.gmail.com:587).
type EmailAuthenticationData struct {
	Server   string
	User     string
	Password string
}

// SendEmail establishes an authenticated session with an external SMTP host
// and iteratively builds/transmits multipart MIME messages containing the requested ebook files.
// Network, authentication, or filesystem IO errors bubble up directly from the underlying libraries.
func SendEmail(s EmailAuthenticationData, msgs ...Email) error {
	serverOnly := strings.Split(s.Server, ":")[0]
	auth := smtp.PlainAuth("", s.User, s.Password, serverOnly)
	for _, msg := range msgs {
		m := email.NewMessage("", "")
		m.To = []string{msg.To}
		m.From = mail.Address{Name: "", Address: s.User}
		for _, file := range msg.Files {
			err := m.Attach(file)
			if err != nil {
				return err
			}
		}
		err := email.Send(s.Server, auth, m)
		if err != nil {
			return err
		}
	}
	return nil
}

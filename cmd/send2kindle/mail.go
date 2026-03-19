package main

import (
	"net/mail"
	"net/smtp"
	"strings"

	"github.com/scorredoira/email"
)

/**
 * Represents the final payload destined for the Kindle device.
 * `Files` should contain fully-qualified paths to the processed book files
 * (post-conversion, if applicable).
 */
type Email struct {
	To    string
	Files []string
}

/**
 * Captures SMTP server details and credentials for dispatching emails.
 * The `Server` field is expected to include the host and port (e.g. "smtp.example.com:587").
 */
type EmailAuthenticationData struct {
	Server   string
	User     string
	Password string
}

/**
 * Orchestrates the dispatch of one or more Email payloads through the provided SMTP relay.
 * This function iterates over messages, attaching multiple books if requested, and blocks
 * until all emails are dispatched or an attachment/transmission error occurs.
 *
 * Side effects:
 * - Reads from disk (for every file listed in the `Files` slice).
 * - Opens a network connection to the defined SMTP server.
 *
 * It terminates early on the first attachment failure or SMTP error, returning it upstream
 * where it is typically expected to be routed through the centralized MustSuccess/Fatalf.
 */
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

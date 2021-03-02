package main

import (
	"net/mail"
	"net/smtp"
	"strings"

	"github.com/scorredoira/email"
)

type Email struct {
    To string
    Files []string
}

type EmailAuthenticationData struct {
    Server string
    User string
    Password string
}

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

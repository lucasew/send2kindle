package main

import (
	"flag"
	"os"
	"path/filepath"
)

// AppConfig Holds all application configuration to avoid global state
type AppConfig struct {
	SmtpUser      string
	SmtpPassword  string
	SmtpServer    string
	KindleEmail   string
	ConvertToEPUB bool
	Books         []string
}

// ParseConfig Creates an AppConfig from flags and environment variables
func ParseConfig() *AppConfig {
	config := &AppConfig{}

	flag.StringVar(&config.SmtpUser, "u", "", "SMTP server username")
	flag.StringVar(&config.SmtpPassword, "p", "", "SMTP server password")
	flag.StringVar(&config.SmtpServer, "s", "", "SMTP server")
	flag.StringVar(&config.KindleEmail, "k", "", "Kindle destination email")
	flag.BoolVar(&config.ConvertToEPUB, "c", false, "Convert to EPUB before sending")
	flag.Parse()

	config.SmtpUser = FallbackStringVariable("SMTP_USER", config.SmtpUser)
	config.SmtpPassword = FallbackStringVariable("SMTP_PASSWD", config.SmtpPassword)
	config.SmtpServer = FallbackStringVariable("SMTP_SERVER", config.SmtpServer)
	config.KindleEmail = FallbackStringVariable("KINDLE_EMAIL", config.KindleEmail)

	args := flag.Args()
	config.Books = make([]string, 0, len(args))
	for _, book := range args {
		abs, err := filepath.Abs(book)
		MustSucess(err)
		_, err = os.Stat(abs)
		MustSucess(err)
		config.Books = append(config.Books, abs)
	}
	if len(config.Books) == 0 {
		Fatalf("No book was specified")
	}

	return config
}

// ProcessAndSend Converts books if necessary and sends the email
func ProcessAndSend(config *AppConfig) error {
	mail_auth := EmailAuthenticationData{
		Server:   config.SmtpServer,
		User:     config.SmtpUser,
		Password: config.SmtpPassword,
	}

	if config.ConvertToEPUB {
		Log("Converting books to EPUB")
		for i := 0; i < len(config.Books); i++ {
			config.Books[i] = ConvertToEPUB(config.Books[i])
		}
	}

	Spew(config.Books)
	Spew(config.KindleEmail)

	email := Email{
		To:    config.KindleEmail,
		Files: config.Books,
	}

	Log("Sending email...")
	return SendEmail(mail_auth, email)
}

func main() {
	defer Cleanup()
	config := ParseConfig()

	MustSucess(ProcessAndSend(config))
	Log("Done!")
}

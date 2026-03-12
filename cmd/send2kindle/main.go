package main

import (
	"flag"
	"os"
	"path/filepath"
)

var ( // Authentication parameters
	SMTP_USER    = ""
	SMTP_PASSWD  = ""
	SMTP_SERVER  = ""
	KINDLE_EMAIL = ""
)

var (
	books         = []string{}
	convertToEPUB = false
)

// init performs early initialization of the CLI layer, handling user-provided flags.
// It supplements missing flags with standard environment variables. Additionally,
// it validates that targeted books actually exist on disk, converting their paths to absolute format.
func init() {
	flag.StringVar(&SMTP_USER, "u", "", "SMTP server username")
	flag.StringVar(&SMTP_PASSWD, "p", "", "SMTP server password")
	flag.StringVar(&SMTP_SERVER, "s", "", "SMTP server")
	flag.StringVar(&KINDLE_EMAIL, "k", "", "Kindle destination email")
	flag.BoolVar(&convertToEPUB, "c", false, "Convert to EPUB before sending")
	flag.Parse()
	SMTP_USER = FallbackStringVariable("SMTP_USER", SMTP_USER)
	SMTP_PASSWD = FallbackStringVariable("SMTP_PASSWD", SMTP_PASSWD)
	SMTP_SERVER = FallbackStringVariable("SMTP_SERVER", SMTP_SERVER)
	KINDLE_EMAIL = FallbackStringVariable("KINDLE_EMAIL", KINDLE_EMAIL)
	args := flag.Args()
	books = make([]string, 0, len(args))
	for _, book := range args {
		abs, err := filepath.Abs(book)
		MustSucess(err)
		_, err = os.Stat(abs)
		MustSucess(err)
		books = append(books, abs)
	}
	if len(books) == 0 {
		Fatalf("No book was specified")
	}
}

// main coordinates the book conversion and dispatch pipeline.
// It guarantees environmental cleanup via defer before orchestrating format conversion (if required)
// and calling the SMTP client layer. Failures bubble up to MustSucess for structured logging.
func main() {
	defer Cleanup()
	mail_auth := EmailAuthenticationData{
		Server:   SMTP_SERVER,
		User:     SMTP_USER,
		Password: SMTP_PASSWD,
	}
	if convertToEPUB {
		Log("Converting books to EPUB")
		for i := 0; i < len(books); i++ {
			books[i] = ConvertToEPUB(books[i])
		}
	}
	Spew(books)
	Spew(KINDLE_EMAIL)
	email := Email{
		To:    KINDLE_EMAIL,
		Files: books,
	}
	Log("Sending email...")
	MustSucess(SendEmail(mail_auth, email))
	Log("Done!")
}

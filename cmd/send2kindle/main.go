package main

import (
	"flag"
	"os"
	"path/filepath"
)

var ( // Authentication parameters
    SMTP_USER = ""
    SMTP_PASSWD = ""
    SMTP_SERVER = ""
    KINDLE_EMAIL = ""
)

var (
    books = []string{}
    convertToMobi = false
)

func init() {
    flag.StringVar(&SMTP_USER, "u", "", "SMTP server username")
    flag.StringVar(&SMTP_PASSWD, "p", "", "SMTP server password")
    flag.StringVar(&SMTP_SERVER, "s", "", "SMTP server")
    flag.StringVar(&KINDLE_EMAIL, "k", "", "Kindle destination email")
    flag.BoolVar(&convertToMobi, "c", false, "Convert to mobi before sending")
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

func main() {
    defer Cleanup()
    mail_auth := EmailAuthenticationData{
        Server: SMTP_SERVER,
        User: SMTP_USER,
        Password: SMTP_PASSWD,
    }
    if (convertToMobi) {
        Log("Converting books to MOBI")
        for i := 0; i < len(books); i++ {
            books[i] = ConvertToMobi(books[i])
        }
    }
    Spew(books)
    Spew(KINDLE_EMAIL)
    email := Email{
        To: KINDLE_EMAIL,
        Files: books,
    }
    Log("Sending email...")
    MustSucess(SendEmail(mail_auth, email))
    Log("Done!")
}

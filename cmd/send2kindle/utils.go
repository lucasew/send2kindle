package main

import (
	"log"
	"os"
	"os/exec"
)

func MustBinary(binary string) {
    _, err := exec.LookPath(binary)
    if err != nil {
        Fatalf("Binary %s not found in $PATH. Aborting...", binary)
    }
}

func MustSucess(err error) {
    if err != nil {
        Fatalf("Fatal error: %s", err)
    }
}

// FallbackStringVariable If string is a empty use environment variable, if env variable is a empty panics
func FallbackStringVariable(env string, def string) string {
    if def != "" {
        return def
    }
    env_value := os.Getenv(env)
    if env_value != "" {
        return env_value
    }
    Fatalf("Neither %s environment variable nor default value is defined", env)
    return "" // dead code just to not give compilation errors
}

func Fatalf(str string, v ...interface{}) {
    log.Fatalf(str, v...)
}

func Log(str string, v ...interface{}) {
    log.Printf(str, v...)
}

func Command(binary string, args ...string) error {
    MustBinary(binary)
    cmd := exec.Command(binary, args...)
    cmd.Stderr = os.Stderr
    cmd.Stdout = os.Stdout
    return cmd.Run()
}

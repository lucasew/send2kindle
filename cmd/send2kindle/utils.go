package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/davecgh/go-spew/spew"
)

// MustBinary ensures that a required system binary is available in the current $PATH.
// It acts as a safety check before executing external commands, terminating the program
// if the binary cannot be resolved.
func MustBinary(binary string) {
	_, err := exec.LookPath(binary)
	if err != nil {
		Fatalf("Binary %s not found in $PATH. Aborting...", binary)
	}
}

// MustSucess acts as the centralized error handler for unrecoverable errors.
// By wrapping operations with this function, the application guarantees that any failure
// is logged with context before forcefully exiting.
func MustSucess(err error) {
	if err != nil {
		Fatalf("Fatal error: %s", err)
	}
}

// FallbackStringVariable resolves a configuration value with a fallback mechanism.
// It first attempts to use the provided default. If empty, it checks the specified environment variable.
// If both are missing, it triggers a fatal error to prevent the application from running with invalid config.
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

// Fatalf logs a formatted error message and immediately terminates the application.
// It serves as the primary exit vector for irrecoverable state failures.
func Fatalf(str string, v ...interface{}) {
	log.Fatalf(str, v...)
}

// Log writes a formatted diagnostic message to standard error.
// Useful for operational visibility during conversion and email dispatch sequences.
func Log(str string, v ...interface{}) {
	log.Printf(str, v...)
}

var showSpewMessages = false

// Spew dumps the deep structure of a given variable if debugging is enabled.
// It provides a richer, type-aware alternative to standard logging for inspecting complex state.
func Spew(v interface{}) {
	if showSpewMessages {
		spew.Dump(v)
	}
}

// Command executes an external binary with the provided arguments, piping stdout and stderr
// directly to the host's standard streams. It implicitly validates the binary's presence
// before attempting execution to avoid opaque system errors.
func Command(binary string, args ...string) error {
	MustBinary(binary)
	cmd := exec.Command(binary, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBuildMessage(t *testing.T) {
	// Create a dummy file to attach
	tmpFile := filepath.Join(os.TempDir(), "dummy_book.epub")
	err := os.WriteFile(tmpFile, []byte("dummy content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}
	defer os.Remove(tmpFile)

	email := Email{
		To:    "kindle@example.com",
		Files: []string{tmpFile},
	}

	from := "sender@example.com"
	msg, err := email.BuildMessage(from)

	if err != nil {
		t.Fatalf("BuildMessage returned unexpected error: %v", err)
	}

	if msg == nil {
		t.Fatalf("BuildMessage returned nil message")
	}

	if len(msg.To) != 1 || msg.To[0] != email.To {
		t.Errorf("Expected To address to be %s, got %v", email.To, msg.To)
	}

	if msg.From.Address != from {
		t.Errorf("Expected From address to be %s, got %s", from, msg.From.Address)
	}
}

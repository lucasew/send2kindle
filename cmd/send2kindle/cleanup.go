package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func CreateTempFileName(extension string) (filename string) {
	tmpdir := os.TempDir()
	generatedName := uuid.New().String()
	filename = filepath.Join(tmpdir, fmt.Sprintf("%s.%s", generatedName, extension))
	AddCleanupHook(func() {
		_ = os.Remove(filename)
	})
	return filename
}

var (
	cleanupHooks = []func(){}
)

func AddCleanupHook(f func()) {
	cleanupHooks = append(cleanupHooks, f)
}

func Cleanup() {
	for _, f := range cleanupHooks {
		f()
	}
}

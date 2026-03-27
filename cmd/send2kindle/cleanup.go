package main

import (
	"fmt"
	"os"
	"path"

	"github.com/google/uuid"
)

func CreateTempFileName(extension string) (filename string) {
	tmpdir := os.TempDir()
	generatedName := uuid.New().String()
	filename = path.Join(tmpdir, fmt.Sprintf("%s.%s", generatedName, extension))
	AddCleanupHook(func() {
		err := os.Remove(filename)
		if err != nil && !os.IsNotExist(err) {
			ReportError(err, "Failed to remove temporary file: "+filename)
		}
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

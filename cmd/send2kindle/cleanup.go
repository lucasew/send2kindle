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
        if err := os.Remove(filename); err != nil && !os.IsNotExist(err) {
            Log("Failed to remove temporary file %s: %v", filename, err)
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

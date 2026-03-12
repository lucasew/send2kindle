package main

import (
	"fmt"
	"os"
	"path"

	"github.com/google/uuid"
)

// CreateTempFileName generates a unique filepath in the system's temporary directory
// with the requested extension. It automatically registers a cleanup hook to guarantee
// the file's deletion once processing finishes, minimizing storage leaks.
func CreateTempFileName(extension string) (filename string) {
	tmpdir := os.TempDir()
	generatedName := uuid.New().String()
	filename = path.Join(tmpdir, fmt.Sprintf("%s.%s", generatedName, extension))
	AddCleanupHook(func() {
		os.Remove(filename)
	})
	return filename
}

var (
	cleanupHooks = []func(){}
)

// AddCleanupHook queues a function to be executed during the teardown phase.
// It is intended to defer temporary file removal and resource deallocation until
// the core email dispatch process concludes.
func AddCleanupHook(f func()) {
	cleanupHooks = append(cleanupHooks, f)
}

// Cleanup iterates over all registered teardown functions and executes them sequentially.
// Designed to be called as a defer at the start of the program's entrypoint, ensuring
// cleanup occurs even if intermediate conversion or network steps fail.
func Cleanup() {
	for _, f := range cleanupHooks {
		f()
	}
}

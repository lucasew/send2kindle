package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func CreateTempFileName(extension string) (filename string) {
	tmpdir, err := ioutil.TempDir("", "send2kindle-*")
	MustSucess(err)
	generatedName := uuid.New().String()
	filename = filepath.Join(tmpdir, fmt.Sprintf("%s.%s", generatedName, extension))
	AddCleanupHook(func() {
		err := os.RemoveAll(tmpdir)
		ReportError(err)
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

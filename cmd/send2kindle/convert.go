package main

// ConvertToEPUB takes an existing file path and invokes the external 'ebook-convert'
// binary to generate an EPUB version of the book. The resulting file is stored in the OS
// temp directory and will be cleaned up automatically during application shutdown.
func ConvertToEPUB(filename string) string {
	outputFileName := CreateTempFileName("epub")
	Log("Converting '%s' to '%s'", filename, outputFileName)
	MustSucess(Command("ebook-convert", filename, outputFileName))
	return outputFileName
}

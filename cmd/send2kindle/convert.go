package main

func ConvertToEPUB(filename string) string {
    outputFileName := CreateTempFileName("epub")
    Log("Converting '%s' to '%s'", filename, outputFileName)
    MustSuccess(Command("ebook-convert", filename, outputFileName))
    return outputFileName
}

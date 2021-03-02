package main

func ConvertToMobi(filename string) string {
    outputFileName := CreateTempFileName("mobi")
    Log("Converting '%s' to '%s'", filename, outputFileName)
    MustSucess(Command("ebook-convert", filename, outputFileName))
    return outputFileName
}

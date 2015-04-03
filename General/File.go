package main

import (
	"os"
)

type CSVFile struct {
	file *os.File
}

func NewCSVFile(fileName string) CSVFile {
	tempFile, _ := os.Create(fileName)
	f := CSVFile{
		file: tempFile}
	return f
}

func (f CSVFile) close() {
	f.file.Close()
}

func (f CSVFile) WriteString(text string) {
	f.file.WriteString(text)
	f.file.Sync()
}

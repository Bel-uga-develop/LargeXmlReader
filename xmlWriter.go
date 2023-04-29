package XmlReader

import (
	"bufio"
	"os"
	"path/filepath"
)

type Writer struct {
	file   *os.File
	writer *bufio.Writer
}

// Create xml file
func (writer *Writer) Create(fileName string) error {
	ext := filepath.Ext(fileName)
	if ext != ".xml" {
		fileName = fileName[0:len(fileName)-len(ext)] + ".xml"
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	writer.file = file
	writer.writer = bufio.NewWriter(writer.file)
	return nil
}

// Close xml file
func (writer *Writer) Close() error {
	if err := writer.file.Close(); err != nil {
		return err
	}
	return nil
}

// Write data
func (writer *Writer) Write(data []byte) error {
	if _, err := writer.writer.Write(data); err != nil {
		return err
	}

	if err := writer.writer.Flush(); err != nil {
		return err
	}

	return nil
}

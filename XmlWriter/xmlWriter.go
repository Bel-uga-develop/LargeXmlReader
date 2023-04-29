// Package for write large xml files line by line
package XmlWriter

import (
	"bufio"
	"os"
)

type Writer struct {
	file *os.File
}

func (writer *Writer) Create(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	writer.file = file

	return nil
}

func (writer *Writer) Close() error {
	if err := writer.file.Close(); err != nil {
		return err
	}
	return nil
}

func (writer *Writer) Write(data []byte) error {
	w := bufio.NewWriter(writer.file)
	w.Write(data)
	w.Flush()
	return nil
}

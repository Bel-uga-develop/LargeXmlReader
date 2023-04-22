package xmlReader

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type Reader struct {
	fileName    string
	elementName string
}

func (reader *Reader) SetFile(fileName string) {
	reader.fileName = fileName
}

func (reader *Reader) SetElement(elementName string) {
	reader.elementName = elementName
}

func (reader *Reader) Read(callBack func(string) error) error {
	file, err := os.Open(reader.fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	if reader.elementName == "" {
		return errors.New("Set value of element: SetElement()")
	}

	scanner := bufio.NewScanner(file)
	element := ""

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "<"+reader.elementName+">") {
			element = scanner.Text()
		} else if strings.Contains(scanner.Text(), "</"+reader.elementName+">") {
			element += scanner.Text()
			err = callBack(element)
			if err != nil {
				return err
			}
			element = ""
		} else if element != "" {
			element += scanner.Text()
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

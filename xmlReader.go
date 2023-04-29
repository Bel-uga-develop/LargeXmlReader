// Package for reading large xml files line by line
package XmlReader

import (
	"bufio"
	"bytes"
	"errors"
	"os"
)

type Reader struct {
	fileName string
	elements []string
}

// Function to specify xml file
func (reader *Reader) SetFile(fileName string) {
	reader.fileName = fileName
}

// Function to specify xml element for loop
func (reader *Reader) SetElement(elementName string) {

	reader.elements = []string{elementName}
}

// Function to specify xml elements for loop
func (reader *Reader) SetElements(elements []string) {
	reader.elements = elements
}

// Reading a xml file line by line
func (reader *Reader) Read(callBack func(string, []byte) error) error {
	file, err := os.Open(reader.fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	if reader.elements == nil {
		return errors.New("set value of element: SetElement()")
	}

	scanner := bufio.NewScanner(file)
	elementSlice := [][]byte{}
	elementName := ""
	for scanner.Scan() {
		row := scanner.Bytes()
		if len(elementSlice) == 0 {
			for _, item := range reader.elements {
				if bytes.Contains(row, []byte("<"+item+">")) {
					elementSlice = append(elementSlice, row)
					elementName = item
					break
				}
			}
		} else if len(elementSlice) != 0 {
			for _, item := range reader.elements {
				if bytes.Contains(row, []byte("</"+item+">")) {
					elementSlice = append(elementSlice, row)
					byteElement := bytes.Join(elementSlice, []byte(""))
					err = callBack(elementName, byteElement)
					if err != nil {
						return err
					}
					elementSlice = [][]byte{}
					elementName = ""
					break
				}
			}
			if len(elementSlice) != 0 {
				elementSlice = append(elementSlice, row)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

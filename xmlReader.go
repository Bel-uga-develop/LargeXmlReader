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
func (reader *Reader) Read(callBack func(string, string) error) error {
	file, err := os.Open(reader.fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	if reader.elements == nil {
		return errors.New("set value of element: SetElement()")
	}

	scanner := bufio.NewScanner(file)
	element := [][]byte{}
	elementName := ""
	for scanner.Scan() {
		row := scanner.Bytes()
		if len(element) == 0 {
			for _, item := range reader.elements {
				if bytes.Contains(row, []byte("<"+item+">")) {
					element = append(element, row)
					elementName = item
					break
				}
			}
		} else if len(element) != 0 {
			for _, item := range reader.elements {
				if bytes.Contains(row, []byte("</"+item+">")) {
					element = append(element, row)
					str := bytes.Join(element, []byte(""))
					err = callBack(elementName, string(str[:]))
					if err != nil {
						return err
					}
					element = [][]byte{}
					elementName = ""
					break
				}
			}
			if len(element) != 0 {
				element = append(element, row)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

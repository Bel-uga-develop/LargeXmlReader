// Package for reading large xml files line by line
package XmlReader

import (
	"bufio"
	"errors"
	"os"
	"strings"
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
		return errors.New("Set value of element: SetElement()")
	}

	scanner := bufio.NewScanner(file)
	element := ""
	elementName := ""
	for scanner.Scan() {
		if element == "" {
			for _, item := range reader.elements {
				if strings.Contains(scanner.Text(), "<"+item+">") {
					element = scanner.Text()
					elementName = item
					break
				}
			}
		} else if element != "" {
			for _, item := range reader.elements {
				if strings.Contains(scanner.Text(), "</"+item+">") {
					element += scanner.Text()
					err = callBack(elementName, element)
					if err != nil {
						return err
					}
					element = ""
					elementName = ""
					break
				}
			}
			if element != "" {
				element += scanner.Text()
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

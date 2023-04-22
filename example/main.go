package main

import (
	"fmt"

	"github.com/Bel-uga-develop/XmlReader"
)

func main() {
	reader := XmlReader.Reader{}
	reader.SetFile("data.xml")
	reader.SetElement("sdnEntry")

	err := reader.Read(readFunc)
	if err != nil {
		fmt.Println(err)
	}
}

func readFunc(element string) error {
	fmt.Println(element)
	fmt.Println("--------------------------------------")
	return nil
}

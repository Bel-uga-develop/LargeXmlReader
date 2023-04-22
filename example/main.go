package main

import (
	"fmt"
)

func main() {
	reader := xmlReader.Reader{}
	reader.SetFile("1.xml")
	//reader.SetElement("sdnEntry")

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

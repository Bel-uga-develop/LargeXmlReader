package main

import (
	"fmt"
	"strconv"

	"github.com/Bel-uga-develop/XmlReader"
)

func main() {
	writer := &XmlReader.Writer{}
	if err := writer.Create("1.xml"); err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 10; i++ {
		data := "<element>\n" + strconv.Itoa(i) + "\n" + "</element>\n"
		if err := writer.Write([]byte(data)); err != nil {
			fmt.Println(err)
		}
	}

	if err := writer.Close(); err != nil {
		fmt.Println(err)
	}
}

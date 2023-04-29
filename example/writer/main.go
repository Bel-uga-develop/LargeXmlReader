package main

import (
	"strconv"

	"github.com/Bel-uga-develop/XmlReader"
)

func main() {
	writer := &XmlReader.Writer{}
	writer.Create("1.xml")
	for i := 0; i < 10; i++ {
		writer.Write([]byte("<element>\n"))
		writer.Write([]byte(strconv.Itoa(i) + "\n"))
		writer.Write([]byte("</element>\n"))
	}
	writer.Close()
}

package main

import (
	"encoding/xml"
	"fmt"

	"github.com/Bel-uga-develop/XmlReader"
)

type PublshInformation struct {
	Publish     string `xml:"Publish"`
	RecordCount int    `xml:"Record_Count"`
}

func main() {
	writer := &XmlReader.Writer{}
	if err := writer.Create("1.xml"); err != nil {
		fmt.Println(err)
	}
	xmldata := &PublshInformation{
		Publish:     "test",
		RecordCount: 1,
	}

	data, _ := xml.MarshalIndent(xmldata, " ", " ")

	if err := writer.Write([]byte(xml.Header)); err != nil {
		fmt.Println(err)
	}

	if err := writer.Write([]byte(data)); err != nil {
		fmt.Println(err)
	}

	if err := writer.Close(); err != nil {
		fmt.Println(err)
	}
}

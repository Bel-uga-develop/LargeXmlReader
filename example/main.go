package main

import (
	"encoding/xml"
	"fmt"

	"github.com/Bel-uga-develop/XmlReader"
)

type SdnEntry struct {
	Uid         int         `xml:"uid"`
	FirstName   string      `xml:"firstName"`
	LastName    string      `xml:"lastName"`
	Title       string      `xml:"title"`
	SdnType     string      `xml:"sdnType"`
	ProgramList ProgramList `xml:"programList"`
	AkaList     AkaList     `xml:"akaList"`
	AddressList AddressList `xml:"addressList"`
}

type ProgramList struct {
	Program []string `xml:"program"`
}

type AkaList struct {
	Aka []Aka `xml:"aka"`
}

type Aka struct {
	Uid      int    `xml:"uid"`
	Type     string `xml:"type"`
	Category string `xml:"category"`
	LastName string `xml:"lastName"`
}

type AddressList struct {
	Address []Address `xml:"address"`
}

type Address struct {
	Uid     int    `xml:"uid"`
	City    string `xml:"city"`
	Country string `xml:"country"`
}

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
	sdnEntry := &SdnEntry{}
	err := xml.Unmarshal([]byte(element), &sdnEntry)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sdnEntry)
	fmt.Println("--------------------------------------")
	return nil
}

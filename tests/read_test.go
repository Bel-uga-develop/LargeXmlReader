package test

import (
	"encoding/xml"
	"github.com/Bel-uga-develop/XmlReader"
	"testing"
	"time"
)

func TestSetFile(t *testing.T) {
	reader := XmlReader.Reader{}
	reader.SetFile("1.xml")
	reader.SetElements([]string{"sdnEntry", "publshInformation"})
	err := reader.Read(readFunc)
	if err == nil {
		t.Error(`error file not exists`)
	}
}

func TestSetElements(t *testing.T) {
	reader := XmlReader.Reader{}
	reader.SetFile("data.xml")
	err := reader.Read(readFunc)
	if err == nil {
		t.Error(`error set elements`)
	}
}

func TestReadElements(t *testing.T) {
	reader := XmlReader.Reader{}
	reader.SetFile("data.xml")
	reader.SetElements([]string{"sdnEntry", "publshInformation"})
	err := reader.Read(readFunc)
	if err != nil {
		t.Error(`error read`)
	}
}

func TestReadElement(t *testing.T) {
	reader := XmlReader.Reader{}
	reader.SetFile("data.xml")
	reader.SetElement("sdnEntry")
	err := reader.Read(readFunc)
	if err != nil {
		t.Error(`error read`)
	}
}

func readFunc(elementName string, element string) error {
	switch elementName {
	case "sdnEntry":
		{
			sdnEntry := &SdnEntry{}
			err := xml.Unmarshal([]byte(element), &sdnEntry)
			if err != nil {
				return err
			}
		}
	case "publshInformation":
		{
			publshInformation := &PublshInformation{}
			err := xml.Unmarshal([]byte(element), &publshInformation)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type userDate time.Time

const userDateFormat = "02/01/2006"

func (ud *userDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	dateString := ""
	err := d.DecodeElement(&dateString, &start)
	if err != nil {
		return err
	}
	dat, err := time.Parse(userDateFormat, dateString)
	if err != nil {
		return err
	}
	*ud = userDate(dat)
	return nil
}

func (ud userDate) String() string {
	return time.Time(ud).Format(time.RFC850)
}

type PublshInformation struct {
	PublishDate userDate `xml:"Publish_Date"`
	RecordCount int      `xml:"Record_Count"`
}

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

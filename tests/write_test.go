package test

import (
	"strconv"
	"testing"

	"github.com/Bel-uga-develop/XmlReader"
)

func TestCreateFile(t *testing.T) {
	writer := &XmlReader.Writer{}
	if err := writer.Create("data_2.xml"); err != nil {
		t.Error(`error create file`)
	}

	for i := 0; i < 10; i++ {
		data := "<element>\n" + strconv.Itoa(i) + "\n" + "</element>\n"
		if err := writer.Write([]byte(data)); err != nil {
			t.Error(`error write file`)
		}
	}

	if err := writer.Close(); err != nil {
		t.Error(`error close file`)
	}
}

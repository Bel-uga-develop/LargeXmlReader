package main

import "strconv"

func main() {
	writer := &Writer{}
	writer.Create("1.xml")
	for i := 0; i < 100000000; i++ {
		writer.Write([]byte("<element>\n"))
		writer.Write([]byte(strconv.Itoa(i) + "\n"))
		writer.Write([]byte("</element>\n"))
	}
	writer.Close()
}

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	bankChannelFilePath := "/Users/taohui/Downloads/渠道.csv"
	// channelFilePath := "/Users/taohui/Downloads/渠道.csv"

	csvFile, _ := os.Open(bankChannelFilePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.LazyQuotes = true

	var contents = []string{}
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		}
		contents = append(contents, "\""+line[0]+"\" to \""+line[1]+"\",")
	}
	fmt.Println(contents)

}

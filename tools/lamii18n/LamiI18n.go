package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("LamiI18n runing")

	csvFile, err := os.Open(`./Lami Live翻译 - V1.0.0.csv`)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	targetFile, err := os.Open("./1.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(line)
		targetFile.Write(line)

	}
	fmt.Println("DONE!!!")
}

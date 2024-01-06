package main

import (
	"fmt"
)

func main() {
	var filePath string

	fmt.Print("Please type csv file path:  ")
	fmt.Scan(&filePath)
	data := readCsvFile(filePath)
	header := getHeader(&data)
	fmt.Println(data)
	fmt.Println("Your file path", header)
}

package main

import (
	"fmt"
)

func main() {
	var filePath string

	fmt.Print("Please type csv file path:  ")
	fmt.Scan(&filePath)
	data := readCsvFile(filePath)
	// header := getHeader(&data)
	getHeader(&data)

	for _, value := range data {
		crwalForsythCounty(value[3])
	}
	// fmt.Println("Your file path", header)
}

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	var filePath string

	fmt.Print("Please type csv file path:  ")
	fmt.Scan(&filePath)
	data := readCsvFile(filePath)

	for i, value := range data {
		fmt.Println(i, value)
	}
	fmt.Println(data)
	fmt.Println("Your file path", filePath)
}
func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return records
}

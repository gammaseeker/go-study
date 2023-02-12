package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello. Please enter the file of expenses to evaluate")

	var filename string
	fmt.Scan(&filename)

	fmt.Println(filename)

	fd, err := os.Open(filename)
	genericErrorHandler(err)

	fmt.Println("Success")

	fileReader := csv.NewReader(fd)
	records, err := fileReader.ReadAll()
	genericErrorHandler(err)
	fmt.Println(records)

	fd.Close()
}

func genericErrorHandler(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

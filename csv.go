package main

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCSV(path string) [][]string {
	csvFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}
	return data

}

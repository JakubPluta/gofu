package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
)

func ReadJSON(path string) map[string]interface{} {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	content, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}

	err = json.Unmarshal([]byte(content), &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

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

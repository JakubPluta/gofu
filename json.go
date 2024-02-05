package main

import (
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

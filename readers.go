package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ListFilesInDir(path string) {
	if path == "" {
		log.Println("No path provided. Using current directory.")
		path = "."
	}
	// List files in a directory
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}

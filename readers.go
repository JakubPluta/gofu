package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func ListFilesInDir(path string) ([]fs.DirEntry, error) {
	if path == "" {
		log.Println("No path provided. Using current directory.")
		path = "."
	}
	// List files in a directory
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func ShowFileInfo(f fs.DirEntry) {
	file, err := f.Info()
	if err != nil {
		return
	}
	fmt.Println(file.Name(), file.Size(), file.IsDir(), file.ModTime())
}

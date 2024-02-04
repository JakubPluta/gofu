package main

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	// CurrDir is the current directory
	CurrDir = "."
	// HomeDir is the home directory
	HomeDir = "~"
	// ParentDir is the parent directory
	ParentDir = ".."
	// RootDir is the root directory
	RootDir = "/"
	// TempDir is the temporary directory
	TempDir = "/tmp"
)

const (
	FileType       = "file"
	DirType        = "dir"
	HiddenFileType = "hidden"
)

type FileInfo struct {
	Name    string
	IsDir   bool
	ModTime string
	Size    int64
}

// GetCurrentWorkingDirectory returns the current working directory.
// No parameters.
// Returns a string and an error.
func GetCurrentWorkingDirectory() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}

// GetHomeDirectory gets the user's home directory.
// Returns a string for the directory path and an error.
func GetHomeDirectory() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return dir, nil
}

// listFiles retrieves a list of files in the specified path, with options to include hidden files and directories.
//
// Parameters:
//
//	path string - the directory path to list files from
//	includeHidden bool - flag to include hidden files
//	includeDirs bool - flag to include directories
//
// Return type(s):
//
//	[]os.DirEntry - a slice of os.DirEntry representing the filtered files
//	error - an error, if any, encountered during the operation
func listFiles(path string, includeHidden, includeDirs bool) ([]os.DirEntry, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var filteredFiles []os.DirEntry

	for _, file := range files {
		if !includeHidden && strings.HasPrefix(file.Name(), ".") {
			continue
		}
		if !includeDirs && file.IsDir() {
			continue
		}
		filteredFiles = append(filteredFiles, file)
	}

	return filteredFiles, nil
}

func GetFilesListRecursively(path string, includeHidden, includeDirs bool) []os.DirEntry {
	var searchResults []os.DirEntry

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		files, err := listFiles(path, includeHidden, includeDirs)
		if err != nil {
			return err
		}
		searchResults = append(searchResults, files...)
		return nil
	})
	return searchResults
}

func GetFileInfo(f os.DirEntry) *FileInfo {
	file, err := f.Info()
	if err != nil {
		return nil
	}
	fileInfo := FileInfo{
		Name:    file.Name(),
		IsDir:   file.IsDir(),
		ModTime: file.ModTime().String(),
		Size:    file.Size(),
	}
	return &fileInfo
}

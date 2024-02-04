package fu

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
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

type FileSizeType int64

func (f FileSizeType) KB() float64 {
	return float64(f) / 1024
}

func (f FileSizeType) MB() float64 {
	return float64(f) / (1024 * 1024)
}

type FileInfo struct {
	Name     string
	FullPath string
	IsDir    bool
	ModTime  time.Time
	Size     FileSizeType
}

// bytesToMegabytes converts bytes to megabytes.
//
// bytes int64
// float64
func bytesToMegabytes(bytes int64) float64 {
	return float64(bytes) / (1024 * 1024)
}

// bytesToKilobytes calculates the number of kilobytes from the given bytes.
//
// bytes int64
// float64
func bytesToKilobytes(bytes int64) float64 {
	return float64(bytes) / 1024
}

// String returns the full path of the file information.
func (f *FileInfo) String() string {
	return fmt.Sprintf("Name: %s, FullPath: %s, IsDir: %t, ModTime: %s, Size: %.2f MB", f.Name, f.FullPath, f.IsDir, f.ModTime.String(), f.Size.MB())
}

// getFullPath returns the full path of the file.
//
// Parameter: fileInfo os.FileInfo
// Return type: string, error
func getFullPath(fileInfo os.FileInfo) (string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(workingDir, fileInfo.Name())
	return fullPath, nil
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

// filterFile filters the file based on the given criteria.
// Parameters: file os.FileInfo, includeHidden bool, includeDirs bool.
// Returns os.FileInfo.
func filterFile(file os.FileInfo, includeHidden, includeDirs bool) os.FileInfo {
	if !includeHidden && strings.HasPrefix(file.Name(), ".") {
		return nil
	}
	if !includeDirs && file.IsDir() {
		return nil
	}
	return file
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
func ListDirectory(path string, includeHidden, includeDirs bool) ([]os.FileInfo, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var filteredFiles []os.FileInfo

	for _, dirEntry := range files {
		f, err := dirEntry.Info()
		if err != nil {
			continue
		}
		file := filterFile(f, includeHidden, includeDirs)
		if file == nil {
			continue
		}

		filteredFiles = append(filteredFiles, file)
	}

	return filteredFiles, nil
}

// GetFilesListRecursively retrieves a list of files recursively from the specified root path.
//
// rootPath string, includeHidden bool, includeDirs bool. []os.FileInfo.
// TODO: improve it
func getFilesListRecursively(rootPath string, includeHidden, includeDirs bool) []os.FileInfo {
	var fileList []os.FileInfo

	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		file := filterFile(info, includeHidden, includeDirs)
		if file == nil {
			return nil
		}
		fileList = append(fileList, file)
		return nil
	})
	return fileList
}

// GetFileInfo returns a FileInfo struct based on the provided os.DirEntry.
//
// entry os.DirEntry
// *FileInfo
func GetFileInfo(entry os.FileInfo) *FileInfo {
	fullPath, err := getFullPath(entry)
	if err != nil {
		fullPath = entry.Name()
	}
	return &FileInfo{
		Name:     entry.Name(),
		IsDir:    entry.IsDir(),
		ModTime:  entry.ModTime(),
		Size:     FileSizeType(entry.Size()),
		FullPath: fullPath,
	}
}

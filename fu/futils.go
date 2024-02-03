package fu

import (
	"os"
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

// GetDirectoryObjectList retrieves the list of directory objects at the specified path.
//
// Parameters:
//
//	path string - the path of the directory
//	includeHidden bool - flag to include hidden files
//
// Return type(s):
//
//	[]os.DirEntry - list of directory objects
//	error - error if any
func GetDirectoryObjectList(path string, includeHidden bool) ([]os.DirEntry, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	if includeHidden {
		return files, nil
	}

	var filteredFiles []os.DirEntry
	for _, file := range files {
		if file.Name()[0] == '.' {
			filteredFiles = append(filteredFiles, file)
		}
	}
	return filteredFiles, nil

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

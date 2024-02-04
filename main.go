package main

import (
	"fmt"

	file_utils "github.com/JakubPluta/gofu/fu/file_utils"
)

func main() {
	f, _ := file_utils.GetFilesListRecursively(".")
	fmt.Println(f)
}

package main

import (
	"fmt"

	fu "github.com/JakubPluta/gofu/fu"
)

func main() {
	f, _ := fu.ListDirectory(".", false, true)
	for _, file := range f {
		fmt.Println(file.Name())
		info := fu.GetFileInfo(file)
		fmt.Println(info)
	}

}

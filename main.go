package main

import (
	"fmt"

	fu "github.com/JakubPluta/gofu/fu"
)

func main() {
	f, _ := fu.ListDirectory(".", false, true)
	for _, file := range f {
		fmt.Println(file.Name())

	}
	sze, _ := fu.GetDirectorySize(".")
	fmt.Println(sze.KB())

	k, _ := fu.ListAllDirectories("../")
	for _, file := range k {
		fmt.Println(file)
	}
}

package main

func main() {
	files, err := ListFilesInDir(".")
	if err != nil {
		return
	}
	for i := 0; i < len(files); i++ {
		ShowFileInfo(files[i])
	}

}

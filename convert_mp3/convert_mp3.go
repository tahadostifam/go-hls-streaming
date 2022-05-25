package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("go run ./convert_mp3.go file.mp3")
		os.Exit(0)
	} else {
		filename := strings.TrimSpace(os.Args[1])
		saveFileName := CleanExtName(filepath.Base(filename), filepath.Ext(filename))

		fmt.Println(saveFileName)
	}
}

func CleanExtName(filename string, ext string) string {
	return filename[0 : len(filename)-len(ext)]
}

package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ListFiles(dir string) {
	f, err := os.Open(dir)
	if err != nil {
		log.Fatalln("Error opening directory: ", err.Error())
	}

	files, err := f.ReadDir(0)
	if err != nil {
		log.Fatalln("Error reading files from dir: ", err.Error())
	}

	for _, v := range files {
		fmt.Printf("%v%v", v.Name(), func() string {
			if v.IsDir() {
				return "/"
			}
			return ""
		}())
		fmt.Println()
		if v.IsDir() {
			newPath := filepath.Join(dir, v.Name())
			fmt.Print("|  ")
			ListFiles(newPath)
		}
	}
}

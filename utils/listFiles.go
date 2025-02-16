package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func ListFiles(dir string, prefix string) {
	f, err := os.Open(dir)
	if err != nil {
		log.Fatalln("Error opening directory: ", err.Error())
	}

	files, err := f.ReadDir(0)
	if err != nil {
		log.Fatalln("Error reading files from dir: ", err.Error())
	}

    sort.Slice(files, func(i, j int) bool {
        return files[i].Name() < files[j].Name()
    })

	for i, v := range files {

        isLast := i == len(files)

        connector := "|--"
        newPrefix := prefix + "|    "

        if isLast {
            connector = "└── "
            newPrefix = prefix + "      "
        }

        fmt.Printf("%s%s%s\n", prefix, connector, v.Name())

		if v.IsDir() {
			newPath := filepath.Join(dir, v.Name())
			ListFiles(newPath, newPrefix)
		}
	}
}

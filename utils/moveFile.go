package utils

import (
	"fmt"
	"os"
)

func MoveFile(sourceFilePath string, destPath string) error {

	err := os.Rename(sourceFilePath, destPath)
	if err != nil {
		return fmt.Errorf("error moving file: %s", err.Error())
	}
	return nil
}

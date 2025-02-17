package utils

import (
	"fmt"
	"os"
)

func DeleteFile(filePath string) error {
	fmt.Println(filePath)
	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("error deleting file: %s", err.Error())
	}
	return nil
}

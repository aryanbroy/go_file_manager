package utils

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

func RenameFile(from string, to string) error {
	slog.Info("Renaming file", slog.String("from", from), slog.String("to", to))
	dirLocation := "/home/aryan/go/dummy_folder"
	fromFileLocation := filepath.Join(dirLocation, from)
	toFileLocation := filepath.Join(dirLocation, to)

	err := os.Rename(fromFileLocation, toFileLocation)
	if err != nil {
		return fmt.Errorf("error renaming file: %v", err.Error())
	}
	return nil
}

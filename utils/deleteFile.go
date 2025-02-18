package utils

import (
	"log/slog"
	"os"
	"sync"
)

func DeleteFile(filePath string, wg *sync.WaitGroup, errChan chan<- error) {
	defer wg.Done()
	err := os.Remove(filePath)	
	if err != nil {
		errChan <- err
		return
	}
	slog.Info("Deleted ", filePath, " successfully")	
}
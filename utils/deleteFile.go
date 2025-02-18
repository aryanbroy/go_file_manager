package utils

import (
	"fmt"
	"os"
	"sync"
)

func DeleteFile(filePath string, wg *sync.WaitGroup, errChan chan<- error) {
	defer wg.Done()
	err := os.Remove(filePath)	
	if err != nil {
		errChan <- err
	}
	fmt.Println("Deleted ", filePath, " successfully")	
}
/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"log/slog"
	"path/filepath"
	"sync"
	"time"

	"github.com/aryanbroy/file_manager/utils"
	"github.com/spf13/cobra"
)

var (
	stringSliceVar []string
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		slog.Info("Deleting file")
		startTime := time.Now()
		
		var wg sync.WaitGroup
		errChan := make(chan error, len(stringSliceVar))

		for _, file := range stringSliceVar {
			wg.Add(1)
			dummyDir := "/home/aryan/go/dummy_folder"
			filePath := filepath.Join(dummyDir, file)

			go utils.DeleteFile(filePath, &wg, errChan)
		}

		go func ()  {
			wg.Wait()
			close(errChan)
		}()

		hasErr := false
		for err := range errChan {
			hasErr = true
			log.Fatalf("Error: %v", err.Error())
		}

		if !hasErr {
			slog.Info("All files deleted successfully")
		}

		slog.Info("Time elapsed","duration",time.Since(startTime))
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringSliceVar(&stringSliceVar, "del", []string{}, "del single or multiple files (eg: --del file1,file2,file3)")
}

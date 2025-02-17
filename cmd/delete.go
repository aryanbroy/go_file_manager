/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"path/filepath"

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

		for _, file := range stringSliceVar {
			dummyDir := "/home/aryan/go/dummy_folder"
			filePath := filepath.Join(dummyDir, file)

			err := utils.DeleteFile(filePath)
			if err != nil {
				log.Fatalln(err.Error())
			}
			fmt.Printf("Deleted %s successfully\n", file)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringSliceVar(&stringSliceVar, "del", []string{}, "del single or multiple files (eg: --del file1,file2,file3)")
}

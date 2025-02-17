/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"log/slog"
	"path/filepath"

	"github.com/aryanbroy/file_manager/utils"
	"github.com/spf13/cobra"
)

// organizeCmd represents the organize command
var organizeCmd = &cobra.Command{
	Use:   "organize",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("Moving file to specified location")
		sourceFile, err := cmd.Flags().GetString("mv")
		if err != nil {
			log.Fatalln("Error getting sourcefile location: ", err.Error())
		}

		destDir, err := cmd.Flags().GetString("to")
		if err != nil {
			log.Fatalln("Error getting destination directory: ", err.Error())
		}

		if sourceFile == "" || destDir == "" {
			log.Fatalln("Source of destination cannot be empty")
		}

		dummyDir := "/home/aryan/go/dummy_folder"

		if err != nil {
			log.Fatalln("Error getting current dir: ", err.Error())
		}

		currentFilePath := filepath.Join(dummyDir, sourceFile)
		destPath := filepath.Join(dummyDir, destDir, sourceFile)

		err = utils.MoveFile(currentFilePath, destPath)

		if err != nil {
			log.Fatalln(err.Error())
		}

		slog.Info("Successfully moved file")
	},
}

func init() {
	rootCmd.AddCommand(organizeCmd)
	organizeCmd.Flags().String("mv", "", "file name you want to move")
	organizeCmd.Flags().String("to", "", "destination directory name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// organizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// organizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

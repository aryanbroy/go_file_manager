/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"log/slog"

	"github.com/aryanbroy/file_manager/utils"
	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		from, err := cmd.Flags().GetString("from")
		if err != nil {
			log.Fatalln("Error getting flag details: ", err.Error())
		}

		to, err := cmd.Flags().GetString("to")
		if err != nil {
			log.Fatalln("Error getting flag details: ", err.Error())
		}

		if from == "" || to == "" {
			log.Fatalln("Empty flag detected, please fill out all flag values")
		}

		err = utils.RenameFile(from, to)
		if err != nil {
			log.Fatalln(err.Error())
		}
		slog.Info("Successfully renamed file")
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)

	renameCmd.PersistentFlags().String("from", "", "old name of the file")
	renameCmd.PersistentFlags().String("to", "", "new name of the file")
	}

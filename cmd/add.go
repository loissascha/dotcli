/*
Copyright © 2024 Sascha Loishandl <sascha.loishandl@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/loissascha/dotcli/internal/project"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [file or folder path(s)]",
	Short: "Add new files/folders to be synced",
	Long: `Add new files or folders to be synced by the sync command

Synced files and folders get copied from their source into the respective folders.
They will get updated every time the sync command gets executed!

You can add multiple files and folders by just adding multiple arguments to the command.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide at least one folder or file to be added")
			return
		}
		project.AddFilesAndFolders(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

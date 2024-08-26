/*
Copyright © 2024 Sascha Loishandl <sascha.loishandl@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a file/folder from the sync",
	Long: `Removes a file or folder from the sync.

Only removes the file/folder from future syncs. Will not touch the files that are already synced to this folder!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This function is not yet implemented. Please remove it by editing the .dotfilescli file!")
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2024 Sascha Loishandl <sascha.loishandl@gmail.com>
*/
package cmd

import (
	"github.com/loissascha/dotcli/internal/project"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Syncs all your files into this folder.",
	Long: `Synchronizes all the files that are marked for sync into this folder.

Make sure you marked all the files/folders with 'dotcli add'!`,
	Run: func(cmd *cobra.Command, args []string) {
		project.RunSync()
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

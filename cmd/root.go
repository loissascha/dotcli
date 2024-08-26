/*
Copyright © 2024 Sascha Loishandl <sascha.loishandl@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dotcli",
	Short: "A cli tool to manage your sync your dotfiles with your git directory",
	Long: `A cli tool to manage and synchronize your dotfiles with your git directory.

Initialize a project folder with 'dotcli init'.
Add files and folders you want to synchronize from your local system with 'dotcli add'.
Synchronize your files/folders from your system into your dotcli project folder with 'dotcli sync'.
This also works in the reverse way with the command 'dotcli restore'. This can be helpful if you reinstalled your system or want to sync multiple devices to the same dotfiles.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dotcli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

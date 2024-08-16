/*
Copyright © 2024 Sascha Loishandl <sascha.loishandl@gmail.com>
*/
package cmd

import (
	"github.com/loissascha/dotcli/internal/project"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a new dotcli synced folder",
	Long: `Initialize a new dotcli project directory.

Inside a dotcli project you can define files and folders from your system which will get synced into the project directory when running the sync command. 
You can create that directory as your git dotfiles project and easily sync all your important dotfiles into your git folder that way.`,
	Run: func(cmd *cobra.Command, args []string) {
		ignoreGitignore, _ := cmd.Flags().GetBool("ignore-gitignore")

		if len(args) == 0 {
			project.InitProject("./", ignoreGitignore)
		} else {
			project.InitProject(args[0], ignoreGitignore)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().Bool("ignore-gitignore", false, "Don't create/update gitignore file")
}

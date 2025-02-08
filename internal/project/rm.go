package project

import (
	"fmt"
	"slices"

	"github.com/loissascha/dotcli/internal/tools"
)

func RemoveFilesAndFolders(input []string) {
	dir := tools.FormatDir("")
	configPath := tools.GetProjectFilePath(dir)
	if !tools.FileExists(configPath) {
		fmt.Println("Project not initialized. Run init command first!")
		return
	}

	files, folders := getProjectConfig(dir)

	files = slices.DeleteFunc(files, func(e string) bool {
		for _, v := range input {
			if v == e {
				fmt.Println("Deleted entry for", v)
				return true
			}
		}
		return false
	})
	folders = slices.DeleteFunc(folders, func(e string) bool {
		for _, v := range input {
			if v == e {
				fmt.Println("Deleted entry for", v)
				return true
			}
		}
		return false
	})
	writeProjectConfig(dir, files, folders)
}

package project

import (
	"fmt"

	"github.com/loissascha/dotcli/internal/tools"
)

func RemoveFilesAndFolders(input []string) {
	dir := tools.FormatDir("")
	configPath := tools.GetProjectFilePath(dir)
	if !tools.FileExists(configPath) {
		fmt.Println("Project not initialized. Run init command first!")
		return
	}
}

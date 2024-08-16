package project

import (
	"fmt"
	"github.com/loissascha/dotcli/internal/tools"
	"os"
)

func AddFilesAndFolders(inputs []string) {

	dir := tools.FormatDir("")
	configPath := tools.GetProjectFilePath(dir)
	if !tools.FileExists(configPath) {
		fmt.Println("Project not initialized. Run init command first!")
		return
	}

	// check all inputs
	// check if it's a valid folder or file

	for _, input := range inputs {
		_, err := os.ReadDir(input)
		if err != nil {
			_, err := os.ReadFile(input)
			if err != nil {
				fmt.Println("Couldn't find file or folder", input)
				return
			}
			addFile(dir, input)
			return
		}
		addFolder(dir, input)
	}

}

func addFolder(dir string, folder string) {
	_, syncFolders := getProjectConfig(dir)
	fexists := false
	for _, f := range syncFolders {
		if f == folder {
			fexists = true
		}
	}

	if !fexists {
		// add new folder to config
		addNewSyncFolder(dir, folder)
	}
}

func addFile(dir string, file string) {
	syncFiles, _ := getProjectConfig(dir)
	fexist := false
	for _, f := range syncFiles {
		if f == file {
			fexist = true
		}
	}

	if !fexist {
		// add new file to config
		addNewSyncFile(dir, file)
	}
}

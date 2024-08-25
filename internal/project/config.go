package project

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/loissascha/dotcli/internal/tools"
)

func addNewSyncFolder(dir string, folder string) {
	dir = tools.FormatDir(dir)
	files, folders := getProjectConfig(dir)
	folders = append(folders, folder)
	writeProjectConfig(dir, files, folders)
}

func addNewSyncFile(dir string, file string) {
	dir = tools.FormatDir(dir)
	files, folders := getProjectConfig(dir)
	files = append(files, file)
	writeProjectConfig(dir, files, folders)
}

func writeProjectConfig(dir string, files []string, folders []string) {
	dir = tools.FormatDir(dir)
	configStr := fmt.Sprintf("[folders]\n")
	for _, f := range folders {
		configStr += fmt.Sprintf("%v\n", f)
	}

	configStr += fmt.Sprintf("[files]\n")
	for _, f := range files {
		configStr += fmt.Sprintf("%v\n", f)
	}

	os.WriteFile(tools.GetProjectFilePath(dir), []byte(configStr), 0755)
}

func getProjectConfig(dir string) (syncFiles []string, syncFolders []string) {
	dir = tools.FormatDir(dir)
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("can't get home dir!")
		return
	}
	filePath := tools.GetProjectFilePath(dir)
	file, err := os.Open(filePath)
	if err != nil {
		panic("Can't read project config file. Does it exist?")
	}
	defer file.Close()

	var files []string
	var folders []string

	scanner := bufio.NewScanner(file)
	readingFiles := false
	readingFolders := false
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		line = strings.TrimSuffix(line, "\n")
		if line == "" {
			continue
		}
		if line == "[folders]" {
			readingFiles = false
			readingFolders = true
			continue
		}
		if line == "[files]" {
			readingFiles = true
			readingFolders = false
			continue
		}

		if strings.Contains(line, "$HOME") {
			line = strings.Replace(line, "$HOME", homeDir, 1)
		}

		if readingFiles {
			files = append(files, line)
		}
		if readingFolders {
			folders = append(folders, line)
		}
	}
	return files, folders
}

package project

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/loissascha/dotcli/internal/tools"
)

func InitProject(dir string, ignoreGitignore bool) {
	if !tools.DirExists(dir) {
		fmt.Println("Directory does not exist!")
		return
	}

	// check if project file exists
	if !tools.FileExists(tools.GetProjectFilePath(dir)) {
		fmt.Println("Project files do not exist. Creating them...")
		createProjectFile(dir)
	} else {
		fmt.Println("Project already initialized")
		return
	}

	// add project file to gitignore file (create if it doesn't exist)
	if !ignoreGitignore {
		gitignoreFilePath := tools.FormatDir(dir) + "/.gitignore"
		if tools.FileExists(gitignoreFilePath) {
			err := updateGitignoreFile(gitignoreFilePath)
			if err != nil {
				fmt.Println("Error reading .gitignore", err)
			} else {
				fmt.Println("Successfully added project files to .gitignore")
			}
		} else {
			err := createGitignoreFile(gitignoreFilePath)
			if err != nil {
				fmt.Println("Error creating .gitignore file!", err)
			} else {
				fmt.Println("Successfully created .gitignore and added project files to it")
			}
		}
	}
}

func createGitignoreFile(fp string) error {
	err := os.WriteFile(fp, []byte(".dotfilescli"), 0755)
	return err
}

func updateGitignoreFile(fp string) error {
	file, err := os.Open(fp)
	if err != nil {
		fmt.Println("Can't read .gitignore file!", err)
		return err
	}
	defer file.Close()

	alreadyInGitignore := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		line = strings.TrimSuffix(line, "\n")
		if line == ".dotfilescli" {
			alreadyInGitignore = true
			break
		}
	}

	if alreadyInGitignore {
		return nil
	}

	gitignoreContents, err := os.ReadFile(fp)
	if err != nil {
	} else {
		newGitignore := fmt.Sprintf("%v\n%v", string(gitignoreContents[:]), ".dotfilescli")
		os.WriteFile(fp, []byte(newGitignore), 0755)
	}
	return err
}

func createProjectFile(dir string) {
	filePath := tools.GetProjectFilePath(dir)
	config := `[folders]

[files]
`
	os.WriteFile(filePath, []byte(config), 0755)
}

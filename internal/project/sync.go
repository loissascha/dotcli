package project

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	// "os/exec"
	"strings"

	"github.com/loissascha/dotcli/internal/tools"
)

func RunSync() {
	dir := tools.FormatDir("")
	configPath := tools.GetProjectFilePath(dir)
	if !tools.FileExists(configPath) {
		fmt.Println("Project not initialized. Run init command first!")
		return
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("can't load home dir.", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	confirm := tools.ReadInput("Are you sure you want to sync your files into this directory? [(y)es/(n)o]", reader)
	if strings.ToLower(confirm) != "yes" && strings.ToLower(confirm) != "y" {
		fmt.Println("Sync cancelled!")
		return
	}

	files, folders := getProjectConfig(dir)

	for _, v := range files {
		file := v
		if strings.Contains(v, homeDir) {
			file = strings.TrimPrefix(file, homeDir)
		}
		file = strings.TrimPrefix(file, "/")

		i := strings.LastIndex(file, "/")
		folderPath := ""
		if i != -1 {
			folderPath = file[:i]
		}

		if folderPath != "" {
			cmd := exec.Command("mkdir", "-p", folderPath)
			_, err := cmd.Output()
			if err != nil {
				fmt.Println("Error mkdir", err)
				return
			}
		}

		cmd := exec.Command("cp", v, file)
		_, err := cmd.Output()
		if err != nil {
			fmt.Println("Error cp file", err)
			return
		}
	}

	for _, v := range folders {
		folder := v
		if strings.Contains(v, homeDir) {
			folder = strings.TrimPrefix(folder, homeDir)
		}
		folder = strings.TrimPrefix(folder, "/")

		cmd := exec.Command("mkdir", "-p", folder)
		_, err := cmd.Output()
		if err != nil {
			fmt.Println("Error mkdir folder", err)
			return
		}

		fromPath := strings.TrimSuffix(v, "/") + "/."

		cmd = exec.Command("cp", "-r", fromPath, folder)
		_, err = cmd.Output()
		if err != nil {
			fmt.Println("Error cp folder", err.Error())
			return
		}
	}

	fmt.Println("Sync successful!")
}

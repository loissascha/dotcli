package project

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/loissascha/dotcli/internal/tools"
)

func RunRestore() {
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
	confirm := tools.ReadInput("Are you sure you want to restore your system files from this directory? This will overwrite all existing system configurations! [(y)es/(n)o]", reader)
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

		i := strings.LastIndex(v, "/")
		folderPath := ""
		if i != -1 {
			folderPath = v[:i]
		}

		// create directory if it doesn't exist
		if folderPath != "" {
			cmd := exec.Command("mkdir", "-p", folderPath)
			_, err := cmd.Output()
			if err != nil {
				fmt.Println("Error mkdir", err)
				return
			}
		}

		fmt.Println("cp", file, v)
		// cmd := exec.Command("cp", file, v)
		// _, err := cmd.Output()
		// if err != nil {
		// 	fmt.Println("Error cp file", err)
		// 	return
		// }
	}

	for _, v := range folders {
		folder := v
		if strings.Contains(v, homeDir) {
			folder = strings.TrimPrefix(folder, homeDir)
		}
		folder = strings.TrimPrefix(folder, "/")
		fromPath := strings.TrimSuffix(v, "/")

		// create directory if it doesn't exist
		cmd := exec.Command("mkdir", "-p", fromPath)
		_, err := cmd.Output()
		if err != nil {
			fmt.Println("Error mkdir folder", err)
			return
		}

		folder = strings.TrimSuffix(folder, "/") + "/."

		fmt.Println("cp -r", folder, fromPath)
		// cmd = exec.Command("cp", "-r", folder, fromPath)
		// _, err = cmd.Output()
		// if err != nil {
		// 	fmt.Println("Error cp folder", err.Error())
		// 	return
		// }
	}
}

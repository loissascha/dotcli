package tools

import (
	"os"
	"strings"
)

func DirExists(dir string) bool {
	dir = FormatDir(dir)
	_, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	return true
}

func FormatDir(dir string) string {
	if strings.TrimSpace(dir) == "" {
		return "."
	}
	return strings.TrimSuffix(dir, "/")
}

func FileExists(filePath string) bool {
	_, err := os.ReadFile(filePath)
	if err != nil {
		return false
	}
	return true
}

func GetProjectFilePath(dir string) string {
	dir = FormatDir(dir)
	dir = dir + "/.dotfilescli"
	return dir
}

package utils

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func GetAbsolutePath(path string) (string, error) {
	return filepath.Abs(path)
}

func DoesPathExist(absPath string) bool {
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		color.Red("Path does not exist: %s", absPath)
		return false
	}
	return true
}

package config

import (
	"os"
	"path/filepath"
	"portal/constants"
)

func ConfigDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return constants.PORTAL_CONFIG_DIR
	}
	return filepath.Join(home, constants.PORTAL_CONFIG_DIR)
}

func ConfigPath() string {
	return filepath.Join(ConfigDir(), constants.PORTAL_CONFIG_FILE)
}

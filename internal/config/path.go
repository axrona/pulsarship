package config

import (
	"os"
	"path/filepath"
	"strings"
)

func ExpandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		home := os.Getenv("HOME")
		path = filepath.Join(home, path[2:])
	}
	return path
}

func GetConfigPath(configFlag string) string {
	if configFlag != "" {
		return ExpandPath(configFlag)
	}

	if envPath := os.Getenv("PULSARSHIP_CONFIG"); envPath != "" {
		return ExpandPath(envPath)
	}

	home := os.Getenv("HOME")
	return filepath.Join(home, ".config", "pulsarship", "pulsarship.toml")
}

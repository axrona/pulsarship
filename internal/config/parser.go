package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	env "github.com/xeyossr/pulsarship/internal"
	"github.com/xeyossr/pulsarship/internal/models"
)

func ExpandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		path = filepath.Join(env.HOME_ENV, path[2:])
	}
	return path
}

func GetConfigPath(configFlag string) string {
	if configFlag != "" {
		return ExpandPath(configFlag)
	}

	if env.PULSARSHIP_CONFIG != "" {
		return ExpandPath(env.PULSARSHIP_CONFIG)
	}

	return filepath.Join(env.HOME_ENV, ".config", "pulsarship", "pulsarship.toml")
}

// Read and parse the configuration file at the given path.
func ParseConfig(file string) (models.PromptConfig, error) {
	var config models.PromptConfig

	if _, err := toml.DecodeFile(file, &config); err != nil {
		return models.PromptConfig{}, fmt.Errorf("%w", err)
	}

	return config, nil
}

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
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

	data, err := os.ReadFile(file)
	if err != nil {
		return models.PromptConfig{}, fmt.Errorf("could not read config file: %w", err)
	}

	if err := toml.Unmarshal(data, &config); err != nil {
		return models.PromptConfig{}, fmt.Errorf("could not parse config file: %w", err)
	}

	return config, nil
}

// Write the default config to the given path.
func WriteDefaultConfig(path string) error {
	data, err := toml.Marshal(DefaultConfig)
	if err != nil {
		return fmt.Errorf("could not encode default config: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return fmt.Errorf("could not create config directory: %w", err)
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("could not write config file: %w", err)
	}

	return nil
}

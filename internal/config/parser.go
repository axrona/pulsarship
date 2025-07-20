package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/xeyossr/pulsarship/internal/models"
)

// Read and parse the configuration file at the given path.
func ParseConfig(file string) (models.PromptConfig, error) {
	var config models.PromptConfig

	if _, err := toml.DecodeFile(file, &config); err != nil {
		return models.PromptConfig{}, fmt.Errorf("%w", err)
	}

	return config, nil
}

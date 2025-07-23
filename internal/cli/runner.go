package cli

import (
	"fmt"
	"io"

	"github.com/xeyossr/pulsarship/internal/components"
	cfg "github.com/xeyossr/pulsarship/internal/config"
)

func RunPrompt(path string, out io.Writer) error {
	configData, err := cfg.ParseConfig(path)
	if err != nil {
		configData = cfg.DefaultConfig
	}

	prompt, err := components.GenPrompt(false, configData)
	if err != nil {
		return fmt.Errorf("could not generate prompt: %w", err)
	}

	fmt.Fprint(out, prompt)
	return nil
}

func RunRightPrompt(path string, out io.Writer) error {
	configData, err := cfg.ParseConfig(path)
	if err != nil {
		configData = cfg.DefaultConfig
	}

	rightPrompt, err := components.GenPrompt(true, configData)
	if err != nil {
		return fmt.Errorf("could not generate right prompt: %w", err)
	}

	fmt.Fprint(out, rightPrompt)
	return nil
}

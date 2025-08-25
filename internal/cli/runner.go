package cli

import (
	"fmt"
	"io"

	"github.com/axrona/pulsarship/internal/components"
	cfg "github.com/axrona/pulsarship/internal/config"
	"github.com/axrona/pulsarship/internal/utils"
)

func RunPrompt(path string, out io.Writer) error {
	configData, err := cfg.ParseConfig(path)
	if err != nil {
		utils.IfNotDebug(func() {
			configData = cfg.DefaultConfig
		}, func() {
			panic(err)
		})
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
		utils.IfNotDebug(func() {
			configData = cfg.DefaultConfig
		}, func() {
			panic(err)
		})
	}

	rightPrompt, err := components.GenPrompt(true, configData)
	if err != nil {
		return fmt.Errorf("could not generate right prompt: %w", err)
	}

	fmt.Fprint(out, rightPrompt)
	return nil
}

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/xeyossr/pulsarship/internal/components"
	cfg "github.com/xeyossr/pulsarship/internal/config"
)

func expandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		home := os.Getenv("HOME")
		path = filepath.Join(home, path[2:])
	}
	return path
}

func Run(path string, out io.Writer) error {
	path = expandPath(path)
	configData, err := cfg.ParseConfig(path)
	if err != nil {
		return fmt.Errorf("could not parse config: %w", err)
	}

	prompt, err := components.GenPrompt(configData)
	if err != nil {
		return fmt.Errorf("could not generate prompt: %w", err)
	}

	fmt.Fprint(out, prompt)
	return nil
}

func main() {
	defaultConfigFile := "pulsarship.toml"
	defaultConfigPath := filepath.Join(os.Getenv("HOME"), ".config", "pulsarship", defaultConfigFile)

	configPath := flag.String("config", defaultConfigPath, "Custom path to pulsarship.toml")
	flag.Parse()

	if err := Run(*configPath, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

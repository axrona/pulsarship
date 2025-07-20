package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xeyossr/pulsarship/internal/components"
	cfg "github.com/xeyossr/pulsarship/internal/config"
	initShell "github.com/xeyossr/pulsarship/internal/init"
)

func expandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		home := os.Getenv("HOME")
		path = filepath.Join(home, path[2:])
	}
	return path
}

func getConfigPath(configFlag string) string {
	if configFlag != "" {
		return expandPath(configFlag)
	}

	if envPath := os.Getenv("PULSARSHIP_CONFIG"); envPath != "" {
		return expandPath(envPath)
	}

	home := os.Getenv("HOME")
	return filepath.Join(home, ".config", "pulsarship", "pulsarship.toml")
}

func Run(path string, out io.Writer) error {
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

var configFlag string

func main() {
	var rootCmd = &cobra.Command{
		Use:   "pulsarship",
		Short: "🚀 The minimal, fast, and customizable shell prompt. ☄🌌️",
		Long:  "Pulsarship is a minimal, fast, and customizable prompt written in Go.",
	}

	rootCmd.PersistentFlags().StringVarP(&configFlag, "config", "c", "", "Path to the config file")

	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Prints the shell function used to execute pulsarship",
	}

	var promptCmd = &cobra.Command{
		Use:   "prompt",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := getConfigPath(configFlag)
			err := Run(path, os.Stdout)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
		},
	}

	var initBashCmd = &cobra.Command{
		Use:   "bash",
		Short: "Prints Bash init script",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(initShell.BashInit())
		},
	}

	var initZshCmd = &cobra.Command{
		Use:   "zsh",
		Short: "Prints Zsh init script",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(initShell.ZshInit())
		},
	}

	var initFishCmd = &cobra.Command{
		Use:   "fish",
		Short: "Prints Fish init script",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(initShell.FishInit())
		},
	}

	initCmd.AddCommand(initBashCmd)
	initCmd.AddCommand(initZshCmd)
	initCmd.AddCommand(initFishCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(promptCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

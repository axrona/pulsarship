package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xeyossr/pulsarship/internal/cli/flags"
	"github.com/xeyossr/pulsarship/internal/config"
	initShell "github.com/xeyossr/pulsarship/internal/init"
)

var RootCmd = &cobra.Command{
	Use:   "pulsarship",
	Short: "🚀 The minimal, fast, and customizable shell prompt ☄🌌️",
	Long:  "🚀🌌️ The minimal, fast, and customizable shell prompt written in Go 🐹",
}
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Prints the shell function used to execute pulsarship",
}

var PromptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		path := config.GetConfigPath(flags.ConfigFlag)
		err := RunPrompt(path, os.Stdout)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	},
}

var RightCmd = &cobra.Command{
	Use:   "right",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		path := config.GetConfigPath(flags.ConfigFlag)
		err := RunRightPrompt(path, os.Stdout)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	},
}

var InitBashCmd = &cobra.Command{
	Use:   "bash",
	Short: "Prints Bash init script",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(initShell.BashInit())
	},
}

var InitZshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Prints Zsh init script",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(initShell.ZshInit())
	},
}

var InitFishCmd = &cobra.Command{
	Use:   "fish",
	Short: "Prints Fish init script",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(initShell.FishInit())
	},
}

package main

import (
	"fmt"
	"os"

	"github.com/xeyossr/pulsarship/internal/cli"
)

func main() {
	cli.RootCmd.PersistentFlags().StringVarP(&cli.ConfigFlag, "config", "c", "", "Path to the config file")

	cli.InitCmd.AddCommand(cli.InitBashCmd)
	cli.InitCmd.AddCommand(cli.InitZshCmd)
	cli.InitCmd.AddCommand(cli.InitFishCmd)
	cli.RootCmd.AddCommand(cli.InitCmd)
	cli.RootCmd.AddCommand(cli.PromptCmd)

	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

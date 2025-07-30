package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xeyossr/pulsarship/internal/cli"
	"github.com/xeyossr/pulsarship/internal/cli/flags"
)

var (
	version   = "dev"
	tag       = "none"
	commit    = "none"
	buildTime = "unknown"
	buildEnv  = "unknown"
)

func printVersion() {
	fmt.Printf(`pulsarship %s
tag:%s
commit_hash:%s
build_time:%s
build_env:%s
`, version, tag, commit, buildTime, buildEnv)
}

func main() {
	cli.RootCmd.PersistentFlags().StringVarP(&flags.ConfigFlag, "config", "c", "", "Path to the config file")
	cli.RootCmd.PersistentFlags().BoolVarP(&flags.ShowVersion, "version", "v", false, "Print version")

	cli.PromptCmd.Flags().BoolVarP(&flags.ShowRight, "right", "r", false, "Print the right prompt instead of left prompt")

	cli.InitCmd.AddCommand(cli.InitBashCmd)
	cli.InitCmd.AddCommand(cli.InitZshCmd)
	cli.InitCmd.AddCommand(cli.InitFishCmd)

	cli.RootCmd.AddCommand(cli.InitCmd)
	cli.RootCmd.AddCommand(cli.PromptCmd)
	cli.RootCmd.AddCommand(cli.GenConfig)

	cli.RootCmd.Run = func(cmd *cobra.Command, args []string) {
		if flags.ShowVersion {
			printVersion()
			os.Exit(0) // Exit after printing version
		}

		if err := cli.RootCmd.Execute(); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	}

	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

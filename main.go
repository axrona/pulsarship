package main

import (
	"fmt"
	"os"

	"github.com/xeyossr/pulsarship/internal/cli"
)

var (
	version   = "dev"
	tag       = "none"
	commit    = "none"
	buildTime = "unknown"
	buildEnv  = "unknown"
)

func printVersion() string {
	return fmt.Sprintf(`pulsarship %s
tag:%s
commit_hash:%s
build_time:%s
build_env:%s
`, version, tag, commit, buildTime, buildEnv)
}

func main() {
	cli.RootCmd.SetVersionTemplate(printVersion())
	cli.RootCmd.Version = printVersion()
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

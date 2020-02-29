package main

import (
	"github.com/spf13/cobra"
	"github.com/usvc/semver/cmd/semver/bump"
)

var (
	cmd *cobra.Command
)

func run(command *cobra.Command, args []string) {
	command.Help()
}

func main() {
	cmd = &cobra.Command{
		Use: "semver",
		Run: run,
	}
	cmd.AddCommand(bump.GetCommand())
	cmd.Execute()
}

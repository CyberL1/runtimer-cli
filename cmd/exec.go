package cmd

import (
	"fmt"
	"github.com/CyberL1/runtimer/api"

	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use: "exec",
	Short: "Runs a standalone runtime",
	Run: execCli,
	DisableFlagParsing: true,
}

func init() {
	rootCmd.AddCommand(execCmd)
}

func execCli(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Print("Provide a runtime")
		return
	}
	run, _ := api.GetRuntime(args[0], "")
	api.ExecuteRuntime(run, args)
}
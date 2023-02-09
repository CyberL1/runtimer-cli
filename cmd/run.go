package cmd

import (
	"fmt"
	"github.com/CyberL1/runtimer/api"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use: "run",
	Short: "Runs a runtime listed in config file",
	Run: run,
	DisableFlagParsing: true,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func run(cmd *cobra.Command, args []string) {
	config, err := api.GetLocalConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	runtime := config.Runtimes[0]
	if len(config.Runtimes) > 1 {
		runtime, err := api.GetPrimaryRuntime(config)
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(args) == 0 {
			api.ExecuteRuntime(runtime, args)
			return
		}

		if args[0] == "-r" {
			runtime, err = api.GetRuntimeFromConfig(args[1])
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	api.ExecuteRuntime(runtime, args)
}
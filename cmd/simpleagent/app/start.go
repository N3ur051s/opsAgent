package app

import (
	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:        "start",
		Deprecated: "Use \"run\" instead to start the Agent",
		RunE:       start,
	}
)

func init() {
	AgentCmd.AddCommand(startCmd)

	startCmd.Flags().StringVarP(&pidfilePath, "pidfile", "p", "", "path to the pidfile")
}

// Start the main loop
func start(cmd *cobra.Command, args []string) error {
	return run(cmd, args)
}

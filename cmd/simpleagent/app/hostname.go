package app

import (
	"context"
	"fmt"
	"simpleagent/conf"

	"simpleagent/pkg/util/hostname"

	"github.com/spf13/cobra"
)

func init() {
	AgentCmd.AddCommand(getHostnameCommand)
}

var getHostnameCommand = &cobra.Command{
	Use:   "hostname",
	Short: "Print the hostname used by the Agent",
	Long:  ``,
	RunE:  doGetHostname,
}

func doGetHostname(cmd *cobra.Command, args []string) error {
	err := conf.LoadConfig(confFilePath)
	if err != nil {
		return fmt.Errorf("unable to set up global agent configuration: %v", err)
	}

	hname, err := hostname.Get(context.TODO())
	if err != nil {
		return fmt.Errorf("Error getting the hostname: %v", err)
	}

	fmt.Println(hname)
	return nil
}

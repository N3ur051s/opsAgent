package app

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	. "opsAgent/pkg/util/exec"
)

func init() {
	AgentCmd.AddCommand(ExecCommand)
}

var ExecCommand = &cobra.Command{
	Use:   "exec",
	Short: "Exec Command used by the Agent",
	Long:  ``,
	RunE:  execCommand,
}

func execCommand(cmd *cobra.Command, args []string) error {
	t := &ExecTask{
		Name: fmt.Sprintf("exec: %s", args[0]),
		Cmd:  args[0],
	}

	out, err := t.Execute(context.TODO())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(out)
	}

	return nil
}

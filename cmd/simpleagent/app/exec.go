package app

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"

	. "simpleagent/pkg/util/exec"
	"simpleagent/pkg/util/task"
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
	c := task.GetDefaultConfig()
	taskPool, err := task.NewTaskPool(c)
	if err != nil {
		fmt.Printf("taskpool init failed: %s", err)
	}
	taskPool.Start()

	defer taskPool.SafeClose()

	for _, cmd := range args {
		t := &ExecTask{
			Name: fmt.Sprintf("exec: %s", cmd),
			Uuid: uuid.New(),
			Cmd:  cmd,
		}

		taskPool.PushTask(t)
	}
	return nil
}

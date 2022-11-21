package exec

import (
	"context"
	"os/exec"
	"strings"

	pb "opsAgent/pkg/proto/pbgo/opsAgent"
	"opsAgent/pkg/util/log"
)

var (
	ResTasks []*pb.ExecTaskReply
)

type ExecTask struct {
	Name string
	Cmd  string
}

func (execTask *ExecTask) Execute(ctx context.Context) (string, error) {
	log.Infof("*Execute* [%s], CMD: %v \n", execTask.Name, execTask.Cmd)
	out, err := exec.CommandContext(ctx, execTask.Cmd).Output()
	if err != nil {
		log.Errorf("error exec: [%s], out: [%s], err: [%s] ", execTask.Cmd, out, err)
		return string(out), err
	}
	return strings.TrimSpace(string(out)), nil
}

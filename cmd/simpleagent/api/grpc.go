package api

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/grpclog"

	pb "simpleagent/pkg/proto/pbgo/simpleagent"
	. "simpleagent/pkg/util/exec"
	"simpleagent/pkg/util/hostname"
	"simpleagent/pkg/util/log"
	"simpleagent/pkg/util/task"

	"simpleagent/pkg/util/grpc"
)

type server struct {
	pb.UnimplementedAgentServer
}

func (s *server) GetHostname(ctx context.Context, in *pb.HostnameRequest) (*pb.HostnameReply, error) {
	h, err := hostname.Get(ctx)
	if err != nil {
		return &pb.HostnameReply{}, err
	}
	return &pb.HostnameReply{Hostname: h}, nil
}

func (s *server) ExecTask(ctx context.Context, in *pb.ExecTasksRequest) (*pb.ExecTasksReply, error) {
	var c = task.GetDefaultConfig()
	var TaskPool, taskpollerr = task.NewTaskPool(c)
	if taskpollerr != nil {
		log.Errorf("taskpool init failed: %s", taskpollerr)
	}
	TaskPool.Start()

	for _, task := range in.Exectasks {
		t := &ExecTask{
			Name: task.Name,
			Uuid: uuid.New(),
			Cmd:  task.Command,
		}
		TaskPool.PushTask(t)
	}

	TaskPool.SafeClose()

	defer FlushRes()

	return &pb.ExecTasksReply{ExecTasksres: ResTasks}, nil
}

func init() {
	grpclog.SetLoggerV2(grpc.NewLogger())
}

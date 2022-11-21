package api

import (
	"context"

	"google.golang.org/grpc/grpclog"

	pb "opsAgent/pkg/proto/pbgo/opsAgent"
	. "opsAgent/pkg/util/exec"
	"opsAgent/pkg/util/hostname"
	. "opsAgent/pkg/util/writefile"

	"opsAgent/pkg/util/grpc"
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

func (s *server) ExecTask(ctx context.Context, in *pb.ExecTaskRequest) (*pb.ExecTaskReply, error) {
	t := &ExecTask{
		Name: in.Name,
		Cmd:  in.Cmd,
	}

	var reserr = ""
	out, err := t.Execute(ctx)
	if err != nil {
		reserr = err.Error()
	}

	result := &pb.ExecTaskReply{
		Name:   in.Name,
		Cmd:    in.Cmd,
		Result: out,
		Error:  reserr,
	}

	return result, nil
}

func (s *server) WriteFile(ctx context.Context, in *pb.WriteFilesRequest) (*pb.WriteFilesReply, error) {
	var wfRess []*pb.WriteFileReply
	for _, file := range in.Files {
		t := &WriteFile{
			Name:    file.Filename,
			Path:    file.Filepath,
			Content: file.Content,
		}
		res, err := t.Execute()
		var wferr = ""
		var code = 200

		if err != nil {
			wferr = err.Error()
			code = 500
		}
		wfRes := &pb.WriteFileReply{
			Filename: file.Filename,
			Error:    wferr,
			Result:   res,
			Code:     int32(code),
		}
		wfRess = append(wfRess, wfRes)
	}

	return &pb.WriteFilesReply{WfRes: wfRess}, nil
}

func init() {
	grpclog.SetLoggerV2(grpc.NewLogger())
}

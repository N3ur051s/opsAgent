# SimpleAgent

## Getting started

To build the Agent you need:
 * [Go](https://golang.org/doc/install) 1.19 or later. You'll also need to set your `$GOPATH` and have `$GOPATH/bin` in your path.
 * CMake version 3.12 or later and a C++ compiler

## Make

Parse dependency files.
```
make tidy
```

Generate amd64 *.gz package
```
make amd64

make linux_amd64.tar.gz
```

## Run

You can run the agent with:
```
./simpleagent run
```

## APIs

SimpleAgent 基于grpc_gateway,可以同时提供 grpc 和 restful风格 API

### api.proto

```
syntax = "proto3";

package simpleagent.api.v1;

import "simpleagent/model.proto";
import "google/api/annotations.proto";

option go_package = "pkg/proto/pbgo"; 


// The hostname service definition.
service Agent {
    rpc GetHostname (simpleagent.model.v1.HostnameRequest) returns (simpleagent.model.v1.HostnameReply) {
        option (google.api.http) = {
            get: "/v1/grpc/host"
        };
    }

    rpc ExecTask (simpleagent.model.v1.ExecTasksRequest) returns (simpleagent.model.v1.ExecTasksReply) {
        option (google.api.http) = {
            post: "/v1/grpc/exec"
            body: "*"
        };
    }
}
```

### model.proto

```
syntax = "proto3";

package simpleagent.model.v1;

option go_package = "pkg/proto/pbgo"; // golang


message HostnameRequest {}

message HostnameReply {
    string hostname = 1;
}

message ExecTaskRequest {
    string name = 1;
    string command = 2;
}

message ExecTasksRequest {
    repeated ExecTaskRequest exectasks = 1;
}

message ExecTaskReply {
    string name = 1;
    string uuid = 2;
    string cmd = 3;
    string result = 4;
    string  error = 5;
    bool success = 6;
}

message ExecTasksReply {
    repeated ExecTaskReply execTasksres = 1;
}

```
#### 服务调用支持TLS加密，目前TLS代码注释掉了，想用的话可以把注释打开


## Usage
支持直接命令行调用
```
$ ./simpleagent

The SimpleAgent faithfully collects events and metrics and brings them
to Server on your behalf so that you can do something useful with your
monitoring and performance data.

Usage:
  ./simpleagent [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  exec        Exec Command used by the Agent
  help        Help about any command
  hostname    Print the hostname used by the Agent
  run         Run the Agent
  stop        Stops a running Agent
  version     Print the version info

Flags:
  -c, --cfgpath string   path to directory containing simpleagent.conf
  -h, --help             help for ./simpleagent
  -n, --no-color         disable color output

Use "./simpleagent [command] --help" for more information about a command.
```
#### 获取hostname
```
$ ./simpleagent hostname
```
#### 异步执行shell任务
```
$ ./simpleagent exec "task1" "task2" "task3"
```

#### 任务配置详解
| 参数                | 类型   | 描述                                                         |
| ------------------- | ------ | ------------------------------------------------------------ |
| MaxTaskNum          | Int    | 单个 任务池 实例能缓存的任务数量上限，默认为 1000。  | |
| MaxIoWorkerNum      | Int64  | 单个 任务池 能并发的最多goroutine的数量，默认为50，该参数可以在代码内根据实际服务器的性能去配置。 |
| MaxRetryTimes       | Int    | 如果某个 任务 首次执行失败，能够对其重试的次数，默认为 10 次。<br/>如果 retries 小于等于 0，该 ProducerBatch 首次发送失败后将直接进入失败队列。 |
| BaseRetryBackOffMs  | Int64  | 首次重试的退避时间，默认为 100 毫秒。 任务池 采用指数退避算法，第 N 次重试的计划等待时间为 baseRetryBackOffMs * 2^(N-1)。 |
| MaxRetryBackOffMs   | Int64  | 重试的最大退避时间，默认为 50 秒。    



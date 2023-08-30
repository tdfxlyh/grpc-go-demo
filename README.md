# 本项目是一个grpc使用的demo，对文件进行了划分。

# 下面是一个grpc简单的使用教程，与本项目关系不大

## 1.go install
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```


## 2.写一个.proto文件
```protobuf
syntax = "proto3";

option go_package = "grpc_gen/calculator";

message AddReq {
  int64 num1 = 1;
  int64 num2 = 2;
}

message AddResp {
  int64 res = 1;
}


service TaskService {
  rpc Add (AddReq) returns (AddResp);
}
```


## 3.使用命令生成代码
```bash
protoc --go_out=. --go-grpc_out=. .\{protoc_dir}\{proto_file}

# 举例
protoc --go_out=. --go-grpc_out=. .\protoc_files\calculator.proto
```


## 4.编写服务端代码
```go
package main

import (
	"context"
	"net"

	"google.golang.org/grpc"
	pb "liuyaohui.lyh/rpcProject/grpc_gen/calculator"
)

type taskServer struct {
	pb.UnimplementedTaskServiceServer // 这里不要忘了
}

func (s *taskServer) Add(ctx context.Context, req *pb.AddReq) (*pb.AddResp, error) {
	myResp := &pb.AddResp{
		Res: req.Num1 + req.Num2,
	}
	return myResp, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTaskServiceServer(grpcServer, &taskServer{})
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
```



## 5.编写客户端代码
```go
package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "liuyaohui.lyh/rpcProject/grpc_gen/calculator"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewTaskServiceClient(conn)

	listResponse, err := client.Add(context.Background(), &pb.AddReq{Num1: 5, Num2: 6})
	if err != nil {
		panic(err)
	}
	log.Printf("Tasks: %v", listResponse.Res)
}
```
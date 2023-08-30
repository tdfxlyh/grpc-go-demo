package main

import (
	pb "github.com/tdfxlyh/grpc-go-demo/grpc_gen/calculator"
	"google.golang.org/grpc"

	"net"
)

func main() {
	InitGRPC()
}

// InitGRPC 作为服务端只实现一个protoc文件
func InitGRPC() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(grpcServer, &CalculatorServer{})
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

package main

import (
	"context"
	pb "github.com/tdfxlyh/grpc-go-demo/grpc_gen/calculator"
	"github.com/tdfxlyh/grpc-go-demo/server/handler"
)

type CalculatorServer struct {
	pb.UnimplementedCalculatorServiceServer
}

func (s *CalculatorServer) Add(ctx context.Context, req *pb.AddReq) (resp *pb.AddResp, err error) {
	h := handler.NewAddHandler(ctx, req)
	h.Process()
	resp, err = h.Resp, h.Err
	return
}

package main

import (
	"context"
	"fmt"
	"github.com/tdfxlyh/grpc-go-demo/client/caller"
	pb "github.com/tdfxlyh/grpc-go-demo/grpc_gen/calculator"
)

func goRpcTest() {
	ctx := context.Background()

	addReq := &pb.AddReq{
		Num1: 6,
		Num2: 0,
	}
	addResp, err := caller.CalculatorServiceClient.Add(ctx, addReq)
	if err != nil {
		fmt.Printf("rpc fail, err=%s\n", err)
	}
	if addResp == nil {
		fmt.Printf("addResp is nil \n")
		return
	}
	fmt.Printf("结果: %d", addResp.Res)
}

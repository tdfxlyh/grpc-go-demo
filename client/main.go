package main

import "github.com/tdfxlyh/grpc-go-demo/client/caller"

func main() {
	caller.InitGRPCClient()

	// 随便调用测试一下
	goRpcTest()
}

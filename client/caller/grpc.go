package caller

import (
	pb "github.com/tdfxlyh/grpc-go-demo/grpc_gen/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	CalculatorServiceClient pb.CalculatorServiceClient
)

func InitGRPCClient() {
	CalculatorServiceClient = NewGRPCClient("localhost:50051")
}

func NewGRPCClient(dial string) pb.CalculatorServiceClient {
	conn, err := grpc.Dial(dial, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	//defer conn.Close() // 不用关
	return pb.NewCalculatorServiceClient(conn)
}

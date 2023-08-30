package handler

import (
	"context"
	"fmt"
	pb "github.com/tdfxlyh/grpc-go-demo/grpc_gen/calculator"
)

type AddHandler struct {
	Ctx context.Context
	Req *pb.AddReq

	Resp *pb.AddResp
	Err  error
}

func NewAddHandler(ctx context.Context, req *pb.AddReq) *AddHandler {
	return &AddHandler{
		Ctx: ctx,
		Req: req,
	}
}

func (h *AddHandler) Process() {
	// 整体捕获一下错误
	defer func() {
		if msg := recover(); msg != nil {
			h.Err = fmt.Errorf("%s", msg)
		}
	}()

	// 校验参数
	// ...
	// 处理数据
	// ...
	// 打包数据
	// ...

	h.Resp = &pb.AddResp{
		Res: h.Req.Num1 + h.Req.Num2,
	}
}

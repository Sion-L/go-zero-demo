package logic

import (
	"context"

	"github.com/Sion-L/go-zero-demo/demo/demo"
	"github.com/Sion-L/go-zero-demo/demo/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DemoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoLogic {
	return &DemoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DemoLogic) Demo(in *demo.DemoReq) (*demo.DemoRes, error) {

	if in.Name == "周志勇" {
		return &demo.DemoRes{
			Message: "傻逼",
		}, nil
	}

	return &demo.DemoRes{
		Message: "周志勇傻逼",
	}, nil
}

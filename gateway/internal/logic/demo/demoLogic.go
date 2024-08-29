package demo

import (
	"context"
	"github.com/Sion-L/go-zero-demo/demo/demo"
	"github.com/Sion-L/go-zero-demo/gateway/internal/svc"
	"github.com/Sion-L/go-zero-demo/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoLogic {
	return &DemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DemoLogic) Demo(req *types.DemoReq) (resp *types.DemoRes, err error) {
	rpcResp, err := l.svcCtx.Demo.Demo(l.ctx, &demo.DemoReq{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &types.DemoRes{
		Message: rpcResp.Message,
	}, nil
}

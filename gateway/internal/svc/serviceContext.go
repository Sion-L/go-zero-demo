package svc

import (
	"github.com/Sion-L/go-zero-demo/demo/mydemo"
	"github.com/Sion-L/go-zero-demo/gateway/internal/config"
	"github.com/Sion-L/go-zero-demo/transform/transformer"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformer.Transformer
	Demo        mydemo.MyDemo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)), // 手动代码
		Demo:        mydemo.NewMyDemo(zrpc.MustNewClient(c.Demo)),
	}
}

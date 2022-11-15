package svc

import (
	"looklook/app/test/cmd/api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/test/cmd/rpc/test"
)

type ServiceContext struct {
	Config config.Config
	TestRpc test.Test
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		TestRpc: test.NewTest(zrpc.MustNewClient(c.TestRpcConf)),
	}
}

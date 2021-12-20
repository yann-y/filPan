package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"ipfsdisk/service/user/api/internal/config"
	"ipfsdisk/service/user/rpc/pb"
	"ipfsdisk/service/user/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc pb.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}

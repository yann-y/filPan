package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"ipfsdisk/service/account/api/internal/config"
	"ipfsdisk/service/account/rpc/pb"
	"ipfsdisk/service/account/rpc/user"
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

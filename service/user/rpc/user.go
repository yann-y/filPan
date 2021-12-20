package main

import (
	"fmt"
	"ipfsdisk/service/user/rpc/internal/config"
	"ipfsdisk/service/user/rpc/internal/server"
	"ipfsdisk/service/user/rpc/internal/svc"
	user "ipfsdisk/service/user/rpc/pb"
	"os"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	//flag.Parse()
	confFilePath, err := os.Getwd()
	if err != nil {
		fmt.Printf("cnfig path err  %s  ", confFilePath)
	}
	confFile := confFilePath + "\\etc\\user.yaml"

	var c config.Config
	conf.MustLoad(confFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

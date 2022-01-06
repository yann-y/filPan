package main

import (
	"fmt"
	"os"
	"path"

	"ipfsdisk/service/account/api/internal/config"
	"ipfsdisk/service/account/api/internal/handler"
	"ipfsdisk/service/account/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

//var configFile = flag.String("f", "etc/pb-api.yaml", "the config file")

func main() {
	//flag.Parse()
	confFilePath, err := os.Getwd()
	if err != nil {
		fmt.Printf("cnfig path err  %s  ", confFilePath)
	}
	confFile := path.Join(confFilePath, "etc/pb-api.yaml")

	var c config.Config
	conf.MustLoad(confFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

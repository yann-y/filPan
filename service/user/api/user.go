package main

import (
	"fmt"
	"os"
	"path"

	"ipfsdisk/service/user/api/internal/config"
	"ipfsdisk/service/user/api/internal/handler"
	"ipfsdisk/service/user/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

//var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	//flag.Parse()
	confFilePath, err := os.Getwd()
	if err != nil {
		fmt.Printf("cnfig path err  %s  ", confFilePath)
	}
	confFile := path.Join(confFilePath, "etc/user-api.yaml")

	var c config.Config
	conf.MustLoad(confFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

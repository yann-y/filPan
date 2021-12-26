package svc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ipfsdisk/service/user/rpc/db/query"
	"ipfsdisk/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel *query.Query
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, _ := gorm.Open(mysql.Open(c.Mysql.DataSource))

	return &ServiceContext{
		Config:    c,
		UserModel: query.Use(db),
	}
}

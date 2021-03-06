// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package logic

import (
	"context"
	"ipfsdisk/service/account/rpc/internal/svc"
	"ipfsdisk/service/account/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserResisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserResisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserResisterLogic {
	return &UserResisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserResister 用户注册
func (l *UserResisterLogic) UserResister(in *user.ResisterReq) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	// 查询用户名或者邮箱是否存在
	// 查看验证码是否有效

	return &user.LoginResponse{}, nil
}

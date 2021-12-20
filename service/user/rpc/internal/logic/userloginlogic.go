package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"

	"ipfsdisk/service/user/rpc/internal/svc"
	"ipfsdisk/service/user/rpc/user"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *user.LoginReq) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line

	return &user.LoginResponse{}, nil
}

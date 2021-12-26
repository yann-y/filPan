package logic

import (
	"context"
	"errors"
	"github.com/tal-tech/go-zero/core/logx"
	"strings"

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
	if len(strings.TrimSpace(in.Username)) == 0 || len(strings.TrimSpace(in.Password)) == 0 {
		return nil, errors.New("用户名或者密码不正确")
	}
	panUser := l.svcCtx.UserModel.PanUser
	userInfo, err := panUser.WithContext(l.ctx).Where(panUser.Username.Eq(in.Username), panUser.Password.Eq(in.Password)).First()
	if err != nil {
		return nil, errors.New("用户名或者密码不正确")
	}
	token := "test"
	return &user.LoginResponse{
		Emile:       userInfo.Email,
		Name:        userInfo.Username,
		AccessToken: token,
	}, nil
}

package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	es "ipfsdisk/pkg/erros"
	"ipfsdisk/pkg/jwt"
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
		return nil, es.PasswdError
	}
	panUser := l.svcCtx.UserModel.PanUser
	userInfo, err := panUser.WithContext(l.ctx).Where(panUser.Username.Eq(in.Username), panUser.Password.Eq(in.Password)).First()
	if err != nil {
		return nil, es.PasswdError
	}
	token, err := jwt.GenerateToken(userInfo.Username, userInfo.Email)
	if err != nil {
		l.Logger.Error("token生成失败--》", err)
	}
	return &user.LoginResponse{
		Emile:       userInfo.Email,
		Name:        userInfo.Username,
		AccessToken: token,
	}, nil
}

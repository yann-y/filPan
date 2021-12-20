// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package user

import (
	"context"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"ipfsdisk/service/user/rpc/pb"
)

type (
	IdRequest     = pb.IdRequest
	LoginReq      = pb.LoginReq
	LoginResponse = pb.UserLoginResponse

	User interface {
		UserLogin(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) UserLogin(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.UserLogin(ctx, in, opts...)
}

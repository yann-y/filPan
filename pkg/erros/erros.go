package errors

import "errors"

var (
	PasswdError = errors.New("用户名或者密码不正确")
)

// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReply struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Gender       string `json:"gender"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type RegisterReq struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChagePasswdReq struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	OldPasswd string `json:"oldPasswd"`
	NewPasswd string `json:"newPasswd"`
}
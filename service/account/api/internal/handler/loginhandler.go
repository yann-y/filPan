package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"ipfsdisk/service/account/api/internal/logic"
	"ipfsdisk/service/account/api/internal/svc"
	"ipfsdisk/service/account/api/internal/types"
)

func loginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), ctx)
		resp, err := l.Login(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"ipfsdisk/service/user/api/internal/logic"
	"ipfsdisk/service/user/api/internal/svc"
	"ipfsdisk/service/user/api/internal/types"
)

func registerHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), ctx)
		resp, err := l.Register(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

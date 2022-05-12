package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"music/service/mv/api/internal/logic"
	"music/service/mv/api/internal/svc"
	"music/service/mv/api/internal/types"
)

func MvFirstHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MvFirstReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMvFirstLogic(r.Context(), svcCtx)
		resp, err := l.MvFirst(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"music/service/dj/api/internal/logic"
	"music/service/dj/api/internal/svc"
	"music/service/dj/api/internal/types"
)

func DjProgramToplistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DjProgramToplistReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDjProgramToplistLogic(r.Context(), svcCtx)
		resp, err := l.DjProgramToplist(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

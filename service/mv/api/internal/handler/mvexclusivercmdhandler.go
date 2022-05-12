package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"music/service/mv/api/internal/logic"
	"music/service/mv/api/internal/svc"
	"music/service/mv/api/internal/types"
)

func MvExclusiveRcmdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MvExclusiveRcmdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMvExclusiveRcmdLogic(r.Context(), svcCtx)
		resp, err := l.MvExclusiveRcmd(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

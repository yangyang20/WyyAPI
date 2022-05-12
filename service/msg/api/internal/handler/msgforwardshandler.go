package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"music/service/msg/api/internal/logic"
	"music/service/msg/api/internal/svc"
	"music/service/msg/api/internal/types"
)

func MsgForwardsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MsgForwardsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMsgForwardsLogic(r.Context(), svcCtx)
		resp, err := l.MsgForwards(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

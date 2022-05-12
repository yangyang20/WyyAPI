package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"music/service/playlist/api/internal/logic"
	"music/service/playlist/api/internal/svc"
	"music/service/playlist/api/internal/types"
)

func PlaylistUpdatePlaycountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PlaylistUpdatePlaycountReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPlaylistUpdatePlaycountLogic(r.Context(), svcCtx)
		resp, err := l.PlaylistUpdatePlaycount(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

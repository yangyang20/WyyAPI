package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"music/service/playlist/api/internal/logic"
	"music/service/playlist/api/internal/svc"
)

func PlaylistHotHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPlaylistHotLogic(r.Context(), svcCtx)
		resp, err := l.PlaylistHot()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

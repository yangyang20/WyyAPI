package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"music/service/song/api/internal/logic"
	"music/service/song/api/internal/svc"
	"music/service/song/api/internal/types"
)

func SongDownloadUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SongDownloadUrlReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSongDownloadUrlLogic(r.Context(), svcCtx)
		resp, err := l.SongDownloadUrl(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

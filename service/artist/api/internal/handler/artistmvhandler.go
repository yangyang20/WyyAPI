package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"music/service/artist/api/internal/logic"
	"music/service/artist/api/internal/svc"
	"music/service/artist/api/internal/types"
)

func ArtistMvHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArtistMvReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewArtistMvLogic(r.Context(), svcCtx)
		resp, err := l.ArtistMv(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

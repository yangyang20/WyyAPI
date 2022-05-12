package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"music/service/artist/api/internal/logic"
	"music/service/artist/api/internal/svc"
	"music/service/artist/api/internal/types"
)

func ArtistAlbumHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArtistAlbumReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewArtistAlbumLogic(r.Context(), svcCtx)
		resp, err := l.ArtistAlbum(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

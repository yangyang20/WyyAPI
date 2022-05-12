package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"music/service/artist/api/internal/logic"
	"music/service/artist/api/internal/svc"
	"music/service/artist/api/internal/types"
)

func ArtistDescHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArtistDescReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewArtistDescLogic(r.Context(), svcCtx)
		resp, err := l.ArtistDesc(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

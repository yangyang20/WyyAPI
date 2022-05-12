package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"music/service/album/api/internal/logic"
	"music/service/album/api/internal/svc"
	"music/service/album/api/internal/types"
)

func AlbumDetailDynamicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AlbumDetailDynamicReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAlbumDetailDynamicLogic(r.Context(), svcCtx)
		resp, err := l.AlbumDetailDynamic(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

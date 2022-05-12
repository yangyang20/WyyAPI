package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"music/service/dj/api/internal/logic"
	"music/service/dj/api/internal/svc"
	"music/service/dj/api/internal/types"
)

func DjRecommendTypeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DjRecommendTypeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDjRecommendTypeLogic(r.Context(), svcCtx)
		resp, err := l.DjRecommendType(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"music/service/search/api/internal/logic"
	"music/service/search/api/internal/svc"
)

func SearchHotDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewSearchHotDetailLogic(r.Context(), svcCtx)
		resp, err := l.SearchHotDetail()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

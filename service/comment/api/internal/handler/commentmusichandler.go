package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"music/service/comment/api/internal/logic"
	"music/service/comment/api/internal/svc"
	"music/service/comment/api/internal/types"
)

func CommentMusicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentMusicReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCommentMusicLogic(r.Context(), svcCtx)
		resp, err := l.CommentMusic(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

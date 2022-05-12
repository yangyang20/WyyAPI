package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"music/service/captcha/api/internal/logic"
	"music/service/captcha/api/internal/svc"
	"music/service/captcha/api/internal/types"
)

func CaptchaVerifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CaptchaVerifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCaptchaVerifyLogic(r.Context(), svcCtx)
		resp, err := l.CaptchaVerify(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

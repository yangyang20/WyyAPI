package logic

import (
	"context"

	"music/service/captcha/api/internal/svc"
	"music/service/captcha/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaVerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) CaptchaVerifyLogic {
	return CaptchaVerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaVerifyLogic) CaptchaVerify(req types.CaptchaVerifyReq) (resp *types.CaptchaVerifyReply, err error) {
	// todo: add your logic here and delete this line

	return
}

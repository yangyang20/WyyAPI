package logic

import (
	"context"

	"music/service/captcha/api/internal/svc"
	"music/service/captcha/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaSentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaSentLogic(ctx context.Context, svcCtx *svc.ServiceContext) CaptchaSentLogic {
	return CaptchaSentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaSentLogic) CaptchaSent(req types.CaptchaSentReq) (resp *types.CaptchaSentReply, err error) {
	// todo: add your logic here and delete this line

	return
}

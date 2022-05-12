package logic

import (
	"context"

	"music/service/login/api/internal/svc"
	"music/service/login/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginQrCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginQrCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginQrCheckLogic {
	return LoginQrCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginQrCheckLogic) LoginQrCheck(req types.LoginQrCheckReq) (resp *types.LoginQrCheckReply, err error) {
	// todo: add your logic here and delete this line

	return
}

package logic

import (
	"context"

	"music/service/login/api/internal/svc"
	"music/service/login/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginQrKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginQrKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginQrKeyLogic {
	return LoginQrKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginQrKeyLogic) LoginQrKey(req types.LoginQrKeyReq) (resp *types.LoginQrKeyReply, err error) {
	// todo: add your logic here and delete this line

	return
}

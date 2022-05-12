package logic

import (
	"context"

	"music/service/login/api/internal/svc"
	"music/service/login/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginQrCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginQrCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginQrCreateLogic {
	return LoginQrCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginQrCreateLogic) LoginQrCreate(req types.LoginQrCreateReq) (resp *types.LoginQrCreateReply, err error) {
	// todo: add your logic here and delete this line

	return
}

package logic

import (
	"context"

	"music/service/login/api/internal/svc"
	"music/service/login/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginStatusLogic {
	return LoginStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginStatusLogic) LoginStatus(req types.LoginStatusReq) (resp *types.LoginStatusReply, err error) {
	// todo: add your logic here and delete this line

	return
}

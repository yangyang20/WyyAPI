package logic

import (
	"context"

	"music/service/login/api/internal/svc"
	"music/service/login/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginRefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginRefreshLogic {
	return LoginRefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginRefreshLogic) LoginRefresh(req types.LoginRefreshReq) (resp *types.LoginRefreshReply, err error) {
	// todo: add your logic here and delete this line

	return
}

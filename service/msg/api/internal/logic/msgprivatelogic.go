package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	. "music/common"
	"music/service/msg/api/internal/svc"
	"music/service/msg/api/internal/types"
)

type MsgPrivateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMsgPrivateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MsgPrivateLogic {
	return MsgPrivateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MsgPrivateLogic) MsgPrivate(req types.MsgPrivateReq) (resp *types.MsgPrivateReply, err error) {
	res := Request("POST", "https://music.163.com/api/msg/private/users", req, Options{
		Crypto: "weapi",
		Cookie: "",
		Proxy:  "",
		RealIP: "",
	})

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err:", err)
	}
	decoder := json.NewDecoder(bytes.NewBuffer(body))
	if err := decoder.Decode(&resp); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return resp, nil
}

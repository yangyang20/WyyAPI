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

type MsgRecentcontactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMsgRecentcontactLogic(ctx context.Context, svcCtx *svc.ServiceContext) MsgRecentcontactLogic {
	return MsgRecentcontactLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MsgRecentcontactLogic) MsgRecentcontact(req types.MsgRecentcontactReq) (resp *types.MsgRecentcontactReply, err error) {
	res := Request("POST", "https://music.163.com/api/msg/recentcontact/get", req, Options{
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

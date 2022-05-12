package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "music/common"
	"music/service/msg/api/internal/svc"
	"music/service/msg/api/internal/types"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type MsgCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMsgCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) MsgCommentsLogic {
	return MsgCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MsgCommentsLogic) MsgComments(req types.MsgCommentsReq) (resp *types.MsgCommentsReply, err error) {
	res := Request("POST", "https://music.163.com/api/v1/user/comments/"+strconv.Itoa(req.Uid), req, Options{
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

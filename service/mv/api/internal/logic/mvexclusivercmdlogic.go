package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "music/common"
	"music/service/mv/api/internal/svc"
	"music/service/mv/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MvExclusiveRcmdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMvExclusiveRcmdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MvExclusiveRcmdLogic {
	return &MvExclusiveRcmdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MvExclusiveRcmdLogic) MvExclusiveRcmd(req *types.MvExclusiveRcmdReq) (resp *types.MvExclusiveRcmdReply, err error) {
	res := Request("POST", "https://interface.music.163.com/api/mv/exclusive/rcmd", req, Options{
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

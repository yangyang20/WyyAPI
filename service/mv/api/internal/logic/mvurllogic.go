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

type MvUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMvUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MvUrlLogic {
	return &MvUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MvUrlLogic) MvUrl(req *types.MvUrlReq) (resp *types.MvUrlReply, err error) {
	res := Request("POST", "https://music.163.com/weapi/song/enhance/play/mv/url", req, Options{
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

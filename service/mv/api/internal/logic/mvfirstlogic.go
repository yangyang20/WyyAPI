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

type MvFirstLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMvFirstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MvFirstLogic {
	return &MvFirstLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MvFirstLogic) MvFirst(req *types.MvFirstReq) (resp *types.MvFirstReply, err error) {
	res := Request("POST", "https://interface.music.163.com/weapi/mv/first", req, Options{
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

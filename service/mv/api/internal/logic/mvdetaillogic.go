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

type MvDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMvDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MvDetailLogic {
	return &MvDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MvDetailLogic) MvDetail(req *types.MvDetailReq) (resp *types.MvDetailReply, err error) {
	res := Request("POST", "https://music.163.com/api/v1/mv/detail", req, Options{
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

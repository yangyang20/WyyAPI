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

type MvDetailInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMvDetailInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MvDetailInfoLogic {
	return &MvDetailInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MvDetailInfoLogic) MvDetailInfo(req *types.MvDetailInfoReq) (resp *types.MvDetailInfoReply, err error) {
	req.Threadid = "R_MV_5_" + req.Threadid
	res := Request("POST", "https://music.163.com/api/comment/commentthread/info", req, Options{
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

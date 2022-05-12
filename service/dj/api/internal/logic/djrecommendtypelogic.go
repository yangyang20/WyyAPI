package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "music/common"
	"music/service/dj/api/internal/svc"
	"music/service/dj/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DjRecommendTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDjRecommendTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) DjRecommendTypeLogic {
	return DjRecommendTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DjRecommendTypeLogic) DjRecommendType(req types.DjRecommendTypeReq) (resp *types.DjRecommendTypeReply, err error) {
	res := Request("POST", "https://music.163.com/weapi/djradio/recommend", req, Options{
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

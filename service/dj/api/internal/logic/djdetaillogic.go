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

type DjDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDjDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) DjDetailLogic {
	return DjDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DjDetailLogic) DjDetail(req types.DjDetailReq) (resp *types.DjDetailReply, err error) {
	res := Request("POST", "https://music.163.com/api/djradio/v2/get", req, Options{
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
	var r *types.DjDetailReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

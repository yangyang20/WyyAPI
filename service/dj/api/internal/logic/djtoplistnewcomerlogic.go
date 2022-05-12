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

type DjToplistNewcomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDjToplistNewcomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) DjToplistNewcomerLogic {
	return DjToplistNewcomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DjToplistNewcomerLogic) DjToplistNewcomer(req types.DjToplistNewcomerReq) (resp *types.DjToplistNewcomerReply, err error) {
	res := Request("POST", "https://music.163.com/api/dj/toplist/newcomer", req, Options{
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

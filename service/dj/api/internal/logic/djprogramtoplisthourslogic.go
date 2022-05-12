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

type DjProgramToplistHoursLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDjProgramToplistHoursLogic(ctx context.Context, svcCtx *svc.ServiceContext) DjProgramToplistHoursLogic {
	return DjProgramToplistHoursLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DjProgramToplistHoursLogic) DjProgramToplistHours(req types.DjProgramToplistHoursReq) (resp *types.DjProgramToplistHoursReply, err error) {
	res := Request("POST", "https://music.163.com/api/djprogram/toplist/hours", req, Options{
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
	var r *types.DjProgramToplistHoursReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

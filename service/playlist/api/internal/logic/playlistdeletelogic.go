package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "music/common"
	"music/service/playlist/api/internal/svc"
	"music/service/playlist/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaylistDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaylistDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) PlaylistDeleteLogic {
	return PlaylistDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaylistDeleteLogic) PlaylistDelete(req types.PlaylistDeleteReq) (resp *types.PlaylistDeleteReply, err error) {
	data := map[string]string{
		"ids": "[" + req.Id + "]",
	}
	res := Request("POST", "https://music.163.com/weapi/playlist/remove", data, Options{
		Crypto: "weapi",
		Cookie: "os=pc",
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

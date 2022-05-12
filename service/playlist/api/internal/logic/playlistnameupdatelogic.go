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

type PlaylistNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaylistNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) PlaylistNameUpdateLogic {
	return PlaylistNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaylistNameUpdateLogic) PlaylistNameUpdate(req types.PlaylistNameUpdateReq) (resp *types.PlaylistNameUpdateReply, err error) {
	res := Request("POST", "https://interface3.music.163.com/eapi/playlist/update/name", req, Options{
		Crypto: "eapi",
		Cookie: "os=pc",
		Proxy:  "",
		RealIP: "",
		Url:    "/api/playlist/update/name",
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

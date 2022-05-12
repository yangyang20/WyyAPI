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

type PlaylistTagsUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaylistTagsUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) PlaylistTagsUpdateLogic {
	return PlaylistTagsUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaylistTagsUpdateLogic) PlaylistTagsUpdate(req types.PlaylistTagsUpdateReq) (resp *types.PlaylistTagsUpdateReqly, err error) {
	res := Request("POST", "https://interface3.music.163.com/eapi/playlist/tags/update", req, Options{
		Crypto: "eapi",
		Cookie: "os=pc",
		Proxy:  "",
		RealIP: "",
		Url:    "/api/playlist/tags/update",
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

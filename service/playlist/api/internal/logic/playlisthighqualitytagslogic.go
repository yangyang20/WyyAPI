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

type PlaylistHighqualityTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaylistHighqualityTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) PlaylistHighqualityTagsLogic {
	return PlaylistHighqualityTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaylistHighqualityTagsLogic) PlaylistHighqualityTags(req types.PlaylistHighqualityTagsReq) (resp *types.PlaylistHighqualityTagsReply, err error) {
	res := Request("POST", "https://music.163.com/api/playlist/highquality/tags", req, Options{
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

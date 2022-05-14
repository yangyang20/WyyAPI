package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "music/common"
	"music/service/song/api/internal/svc"
	"music/service/song/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SongDownloadUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSongDownloadUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) SongDownloadUrlLogic {
	return SongDownloadUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SongDownloadUrlLogic) SongDownloadUrl(req types.SongDownloadUrlReq) (resp *types.SongDownloadUrlReply, err error) {
	res := Request("POST", "https://interface.music.163.com/eapi/song/enhance/download/url", req, Options{
		Crypto: "eapi",
		Cookie: "",
		Proxy:  "",
		RealIP: "",
		Url:    "/api/song/enhance/download/url",
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

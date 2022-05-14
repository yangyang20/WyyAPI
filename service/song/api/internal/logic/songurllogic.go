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

type SongUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSongUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) SongUrlLogic {
	return SongUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SongUrlLogic) SongUrl(req types.SongUrlReq) (resp *types.SongUrlReply, err error) {

	data := map[string]string{
		"ids": "[" + string(req.Id) + "]",
		"br":  string(req.Br),
	}

	res := Request("POST", "https://interface3.music.163.com/eapi/song/enhance/player/url", data, Options{
		Crypto: "eapi",
		Cookie: "os=pc",
		Proxy:  "",
		RealIP: "",
		Url:    "/api/song/enhance/player/url",
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

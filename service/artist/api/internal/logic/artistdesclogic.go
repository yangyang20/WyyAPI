package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "music/common"
	"music/service/artist/api/internal/svc"
	"music/service/artist/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArtistDescLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArtistDescLogic(ctx context.Context, svcCtx *svc.ServiceContext) ArtistDescLogic {
	return ArtistDescLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArtistDescLogic) ArtistDesc(req types.ArtistDescReq) (resp *types.ArtistDescReply, err error) {
	res := Request("POST", "https://music.163.com/weapi/artist/introduction", req, Options{
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
	var r *types.ArtistDescReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

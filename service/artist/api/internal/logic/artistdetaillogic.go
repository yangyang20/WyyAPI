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

type ArtistDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArtistDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) ArtistDetailLogic {
	return ArtistDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArtistDetailLogic) ArtistDetail(req types.ArtistDetailReq) (resp *types.ArtistDetailReply, err error) {
	res := Request("POST", "https://music.163.com/api/artist/head/info/get", req, Options{
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
	var r *types.ArtistDetailReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

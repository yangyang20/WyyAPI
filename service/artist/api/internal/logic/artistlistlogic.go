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

type ArtistListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArtistListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ArtistListLogic {
	return ArtistListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArtistListLogic) ArtistList(req types.ArtistListReq) (resp *types.ArtistListReply, err error) {
	data := map[string]interface{}{
		"initial": "",
		"offset":  req.Offset,
		"limit":   req.Limit,
		"total":   true,
		"type":    req.Type,
		"area":    req.Area,
	}
	res := Request("POST", "https://music.163.com/api/v1/artist/list", data, Options{
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
	var r *types.ArtistListReply
	if err := decoder.Decode(&r); err != nil {
		panic(err)
	}

	return r, nil
}

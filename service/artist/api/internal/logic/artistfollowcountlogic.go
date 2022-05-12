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

type ArtistFollowCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArtistFollowCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) ArtistFollowCountLogic {
	return ArtistFollowCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArtistFollowCountLogic) ArtistFollowCount(req types.ArtistFollowCountReq) (resp *types.ArtistFollowCountReply, err error) {
	res := Request("POST", "https://music.163.com/weapi/artist/follow/count/get", req, Options{
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
	var r *types.ArtistFollowCountReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

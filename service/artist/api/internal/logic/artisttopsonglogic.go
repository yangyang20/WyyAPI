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

type ArtistTopSongLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArtistTopSongLogic(ctx context.Context, svcCtx *svc.ServiceContext) ArtistTopSongLogic {
	return ArtistTopSongLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArtistTopSongLogic) ArtistTopSong(req types.ArtistTopSongReq) (resp *types.ArtistTopSongReply, err error) {
	res := Request("POST", "https://music.163.com/api/artist/top/song", req, Options{
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
	var r *types.ArtistTopSongReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

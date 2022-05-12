package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "music/common"
	"music/service/album/api/internal/svc"
	"music/service/album/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlbumNewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlbumNewLogic(ctx context.Context, svcCtx *svc.ServiceContext) AlbumNewLogic {
	return AlbumNewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlbumNewLogic) AlbumNew(req types.AlbumNewReq) (resp *types.AlbumNewReply, err error) {
	res := Request("POST", "https://music.163.com/weapi/album/new", req, Options{
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
	var r *types.AlbumNewReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

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

type AlbumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) AlbumLogic {
	return AlbumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlbumLogic) Album(req types.AlbumReq) (resp *types.AlbumReply, err error) {
	res := Request("POST", "https://music.163.com/weapi/v1/album/"+string(req.Id), req, Options{
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
	var r *types.AlbumReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

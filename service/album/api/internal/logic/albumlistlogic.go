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

type AlbumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) AlbumListLogic {
	return AlbumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlbumListLogic) AlbumList(req types.AlbumListReq) (resp *types.AlbumListReply, err error) {
	res := Request("POST", "https://music.163.com/weapi/vipmall/albumproduct/list", req, Options{
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
	var r *types.AlbumListReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

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

type AlbumSongsaleboardLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlbumSongsaleboardLogic(ctx context.Context, svcCtx *svc.ServiceContext) AlbumSongsaleboardLogic {
	return AlbumSongsaleboardLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlbumSongsaleboardLogic) AlbumSongsaleboard(req types.AlbumSongsaleboardReq) (resp *types.AlbumSongsaleboardReply, err error) {
	var data map[string]interface{}
	if req.Type == "year" {
		data = map[string]interface{}{
			"albumType": req.AlbumType,
			"year":      req.Year,
			"type":      req.Type,
		}
	} else {
		data = map[string]interface{}{
			"AlbumType": req.Type,
			"type":      req.Type,
		}
	}
	res := Request("POST", "https://music.163.com/api/feealbum/songsaleboard/"+req.Type+"/type", data, Options{
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
	var r *types.AlbumSongsaleboardReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

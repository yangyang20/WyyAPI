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
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type SongDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSongDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) SongDetailLogic {
	return SongDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SongDetailLogic) SongDetail(req types.SongDetailReq) (resp *types.SongDetailReply, err error) {
	strA := strings.Split(req.Ids, "/\\s*,\\s*/")
	var ids string
	for _, v := range strA {
		ids += "{\"id\":" + v + "}"
	}

	data := map[string]string{
		"c": "[" + ids + "]",
	}

	res := Request("POST", "https://music.163.com/api/v3/song/detail", data, Options{
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
	if err := decoder.Decode(&resp); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return resp, nil
}

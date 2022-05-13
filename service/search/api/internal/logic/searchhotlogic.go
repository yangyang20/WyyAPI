package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "music/common"
	"music/service/search/api/internal/svc"
	"music/service/search/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchHotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchHotLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchHotLogic {
	return SearchHotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchHotLogic) SearchHot(req types.SearchHotReq) (resp *types.SearchHotReply, err error) {
	res := Request("POST", "https://music.163.com/weapi/search/hot", req, Options{
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

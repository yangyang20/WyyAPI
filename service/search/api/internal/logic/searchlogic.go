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

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchLogic {
	return SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req types.SearchReq) (resp *types.SearchReply, err error) {
	var data map[string]interface{}
	var url string
	if req.Type == "2000" {
		data = map[string]interface{}{
			"keyword": req.Keywords,
			"scene":   "normal",
			"limit":   req.Limit,
			"offset":  req.Offset,
		}
		url = "https://music.163.com/api/search/voice/get"
	} else {
		data = map[string]interface{}{
			"s":      req.Keywords,
			"type":   req.Type,
			"limit":  req.Limit,
			"offset": req.Offset,
		}
		url = "https://music.163.com/weapi/search/get"
	}
	res := Request("POST", url, data, Options{
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

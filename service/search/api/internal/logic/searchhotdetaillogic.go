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

type SearchHotDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchHotDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchHotDetailLogic {
	return SearchHotDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchHotDetailLogic) SearchHotDetail() (resp *types.SearchHotDetailReply, err error) {
	res := Request("POST", "https://music.163.com/weapi/hotsearchlist/get", nil, Options{
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

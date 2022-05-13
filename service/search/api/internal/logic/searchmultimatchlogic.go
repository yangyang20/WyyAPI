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

type SearchMultimatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchMultimatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchMultimatchLogic {
	return SearchMultimatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchMultimatchLogic) SearchMultimatch(req types.SearchMultimatchReq) (resp *types.SearchMultimatchReply, err error) {
	res := Request("POST", "https://music.163.com/weapi/search/suggest/multimatch", req, Options{
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

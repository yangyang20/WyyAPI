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

type SearchDefaultLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchDefaultLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchDefaultLogic {
	return SearchDefaultLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchDefaultLogic) SearchDefault(req types.SearchDefaultReq) (resp *types.SearchDefaultReply, err error) {
	res := Request("POST", "https://interface3.music.163.com/eapi/search/defaultkeyword/get", req, Options{
		Crypto: "eapi",
		Cookie: "",
		Proxy:  "",
		RealIP: "",
		Url:    "/api/search/defaultkeyword/get",
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

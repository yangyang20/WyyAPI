package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "music/common"
	"music/service/dj/api/internal/svc"
	"music/service/dj/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DjCategoryExcludehotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDjCategoryExcludehotLogic(ctx context.Context, svcCtx *svc.ServiceContext) DjCategoryExcludehotLogic {
	return DjCategoryExcludehotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DjCategoryExcludehotLogic) DjCategoryExcludehot(req types.DjCategoryExcludehotReq) (resp *types.DjCategoryExcludehotReqly, err error) {
	res := Request("POST", "https://music.163.com/weapi/djradio/category/excludehot", req, Options{
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
	var r *types.DjCategoryExcludehotReqly
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

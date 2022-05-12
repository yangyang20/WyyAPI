package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	. "music/common"
	"music/service/mv/api/internal/svc"
	"music/service/mv/api/internal/types"
)

type MvAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMvAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MvAllLogic {
	return &MvAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MvAllLogic) MvAll(req *types.MvAllReq) (resp *types.MvAllReply, err error) {

	data := map[string]interface{}{
		"tags": json.Marshal(map[string]string{
			"地区": req.Area,
			"类型": req.Type,
			"排序": req.Order,
		}),
		"offset": req.Offset,
		"total":  req.Total,
		"limit":  req.Limit,
	}

	res := Request("POST", "https://interface.music.163.com/api/mv/all", data, Options{
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

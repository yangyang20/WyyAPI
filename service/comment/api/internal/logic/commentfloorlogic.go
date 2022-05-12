package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "music/common"
	"music/service/comment/api/internal/config"
	"music/service/comment/api/internal/svc"
	"music/service/comment/api/internal/types"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentFloorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentFloorLogic(ctx context.Context, svcCtx *svc.ServiceContext) CommentFloorLogic {
	return CommentFloorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentFloorLogic) CommentFloor(req types.CommentFloorReq) (resp *types.CommentFloorReply, err error) {
	req.Type = config.ResourceTypeMap(req.Type)
	data := map[string]interface{}{
		"parentCommentId": req.ParentCommentId,
		"threadId":        req.Type + strconv.Itoa(req.Id),
		"time":            req.Time,
		"limit":           req.Limit,
	}
	res := Request("POST", "https://music.163.com/api/resource/comment/floor/get", data, Options{
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
	var r *types.CommentFloorReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

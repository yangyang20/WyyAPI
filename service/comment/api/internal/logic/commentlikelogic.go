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

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) CommentLikeLogic {
	return CommentLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentLikeLogic) CommentLike(req types.CommentLikeReq) (resp *types.CommentLikeReply, err error) {

	req.Type = config.ResourceTypeMap(req.Type)
	data := map[string]interface{}{
		"threadId":  req.Type + req.Id,
		"commentId": req.Cid,
	}
	if req.Type == "A_EV_2_" {
		data["threadId"] = req.ThreadId
	}
	res := Request("POST", "https://music.163.com/weapi/v1/comment/"+req.T, req, Options{
		Crypto: "weapi",
		Cookie: "os=pc",
		Proxy:  "",
		RealIP: "",
	})

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err:", err)
	}
	decoder := json.NewDecoder(bytes.NewBuffer(body))
	var r *types.CommentLikeReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

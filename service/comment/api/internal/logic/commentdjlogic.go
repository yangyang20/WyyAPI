package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "music/common"
	"music/service/comment/api/internal/svc"
	"music/service/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentDjLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentDjLogic(ctx context.Context, svcCtx *svc.ServiceContext) CommentDjLogic {
	return CommentDjLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentDjLogic) CommentDj(req types.CommentDjReq) (resp *types.CommentDjReply, err error) {
	res := Request("POST", "https://music.163.com/weapi/v1/resource/comments/A_DJ_1_"+req.Rid, req, Options{
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
	var r *types.CommentDjReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

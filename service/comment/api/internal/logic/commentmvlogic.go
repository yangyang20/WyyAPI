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

type CommentMvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentMvLogic(ctx context.Context, svcCtx *svc.ServiceContext) CommentMvLogic {
	return CommentMvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentMvLogic) CommentMv(req types.CommentMvReq) (resp *types.CommentMvReply, err error) {
	res := Request("POST", "https://music.163.com/weapi/v1/resource/comments/R_MV_5_"+req.Rid, req, Options{
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
	var r *types.CommentMvReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

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

type CommentPlayListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentPlayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) CommentPlayListLogic {
	return CommentPlayListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentPlayListLogic) CommentPlayList(req types.CommentPlayListReq) (resp *types.CommentPlayListReply, err error) {
	res := Request("POST", "https://music.163.com/weapi/v1/resource/comments/A_PL_0_"+req.Rid, req, Options{
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
	var r *types.CommentPlayListReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

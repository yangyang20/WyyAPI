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

type CommentMusicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentMusicLogic(ctx context.Context, svcCtx *svc.ServiceContext) CommentMusicLogic {
	return CommentMusicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentMusicLogic) CommentMusic(req types.CommentMusicReq) (resp *types.CommentMusicReply, err error) {
	res := Request("POST", "https://music.163.com/api/v1/resource/comments/R_SO_4_"+req.Rid, req, Options{
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
	var r *types.CommentMusicReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

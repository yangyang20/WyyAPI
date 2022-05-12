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

type CommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) CommentLogic {
	return CommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentLogic) Comment(req types.CommentReq) (resp *types.CommentReply, err error) {

	req.T = map[string]string{
		"1": "add",
		"0": "delete",
		"2": "reply",
	}[req.T]
	req.Type = config.ResourceTypeMap(req.Type)

	data := map[string]interface{}{
		"threadId": req.Type + strconv.Itoa(req.Id),
	}

	if (req.Type == "A_EV_2_") {
		data["threadId"] = req.ThreadId
	}
	if (req.T == "add") {
		data["content"] = req.Content
	} else if (req.T == "delete") {
		data["commentId"] = req.CommentId
	} else if (req.T == "reply") {
		data["commentId"] = req.CommentId
		data["content"] = req.Content
	}

	res := Request("POST", "https://music.163.com/weapi/resource/comments/"+req.T, data, Options{
		Crypto: "weapi",
		Cookie: "os=android",
		Proxy:  "",
		RealIP: "",
	})

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err:", err)
	}
	decoder := json.NewDecoder(bytes.NewBuffer(body))
	var r *types.CommentReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

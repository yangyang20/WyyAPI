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

type CommentHugListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentHugListLogic(ctx context.Context, svcCtx *svc.ServiceContext) CommentHugListLogic {
	return CommentHugListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentHugListLogic) CommentHugList(req types.CommentHugListReq) (resp *types.CommentHugListReply, err error) {
	req.Type = config.ResourceTypeMap(req.Type)
	threadId := req.Type + req.Sid
	data := map[string]interface{}{
		"targetUserId": req.Uid,
		"commentId":    req.Cid,
		"cursor":       req.Cursor,
		"threadId":     threadId,
		"pageNo":       req.Page,
		"idCursor":     req.IdCursor,
		"pageSize":     req.PageSize,
	}

	res := Request("POST", "https://music.163.com/api/v2/resource/comments/hug/list", data, Options{
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
	var r *types.CommentHugListReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

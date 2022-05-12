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

type CommentNewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentNewLogic(ctx context.Context, svcCtx *svc.ServiceContext) CommentNewLogic {
	return CommentNewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentNewLogic) CommentNew(req types.CommentNewRep) (resp *types.CommentNewReply, err error) {

	req.Type = config.ResourceTypeMap(req.Type)
	threadId := req.Type + req.Id

	if req.SortType == 1 {
		req.SortType = 99
	}
	var cursor string
	switch req.SortType {
	case 99:
		cursor = strconv.Itoa((req.PageNo - 1) * req.PageSize)
		break
	case 2:
		cursor = "normalHot#" + strconv.Itoa((req.PageNo-1)*req.PageSize)
		break
	case 3:
		cursor = req.Cursor
		break
	default:
		break
	}

	data := map[string]interface{}{
		"threadId":  threadId,
		"pageNo":    req.PageNo,
		"pageSize":  req.PageSize,
		"showInner": req.ShowInner,
		"cursor":    cursor,
		"sortType":  req.SortType,
	}

	res := Request("POST", "https://music.163.com/api/v2/resource/comments", data, Options{
		Crypto: "eapi",
		Cookie: "os=pc",
		Proxy:  "",
		RealIP: "",
		Url:    "/api/v2/resource/comments",
	})

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err:", err)
	}
	decoder := json.NewDecoder(bytes.NewBuffer(body))
	var r *types.CommentNewReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

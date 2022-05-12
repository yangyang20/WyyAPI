package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "music/common"
	"music/service/user/api/internal/svc"
	"music/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserAccountLogic {
	return UserAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAccountLogic) UserAccount(req types.UserAccountReq) (resp *types.UserAccountReply, err error) {

	res := Request("POST", "https://music.163.com/api/nuser/account/get", req, Options{
		Crypto: "weapi",
		Cookie: "",
		Proxy:  "",
		RealIP: "",
		Ua:     "pc",
	})
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err:", err)
	}
	decoder := json.NewDecoder(bytes.NewBuffer(body))
	var r *types.UserAccountReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	return r, nil
}

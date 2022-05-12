package logic

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "music/common"
	"music/service/login/api/internal/svc"
	"music/service/login/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginCellphoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginCellphoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginCellphoneLogic {
	return LoginCellphoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginCellphoneLogic) LoginCellphone(req types.LoginCellphoneReq) (resp *types.LoginCellphoneReply, cookie []*http.Cookie, err error) {
	m := md5.New()
	m.Write([]byte(req.Password))
	req.Password = hex.EncodeToString(m.Sum(nil))
	res := Request("POST", "https://music.163.com/api/login/cellphone", req, Options{
		Crypto: "weapi",
		Cookie: "os=pc; appver=2.9.7",
		Proxy:  "",
		RealIP: "",
		Ua:     "pc",
	})
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err:", err)
		//panic(err)
	}
	decoder := json.NewDecoder(bytes.NewBuffer(body))
	var r *types.LoginCellphoneReply
	if err := decoder.Decode(&r); err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	var cookies string
	for _, v := range res.Cookies() {
		cookies += v.Name + "=" + v.Value + ";"
	}
	r.Cookie = cookies
	return r, res.Cookies(), nil
}

// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"music/service/song/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/song/detail",
				Handler: SongDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/song/download/url",
				Handler: SongDownloadUrlHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/song/url",
				Handler: SongUrlHandler(serverCtx),
			},
		},
	)
}

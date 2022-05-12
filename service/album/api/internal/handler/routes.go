// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"music/service/album/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/album",
				Handler: AlbumHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/album/detail/dynamic",
				Handler: AlbumDetailDynamicHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/album/list",
				Handler: AlbumListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/album/list/style",
				Handler: AlbumListStyleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/album/new",
				Handler: AlbumNewHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/album/newest",
				Handler: AlbumNewestHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/album/songsaleboard",
				Handler: AlbumSongsaleboardHandler(serverCtx),
			},
		},
	)
}

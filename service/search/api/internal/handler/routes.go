// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"music/service/search/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/search",
				Handler: SearchHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/search/default",
				Handler: SearchDefaultHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/search/hot",
				Handler: SearchHotHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/search/hot/detail",
				Handler: SearchHotDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/search/multimatch",
				Handler: SearchMultimatchHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/search/suggest",
				Handler: SearchSuggestHandler(serverCtx),
			},
		},
	)
}
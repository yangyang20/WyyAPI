// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"music/service/msg/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/msg/comments",
				Handler: MsgCommentsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/msg/forwards",
				Handler: MsgForwardsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/msg/notices",
				Handler: MsgNoticesHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/msg/private",
				Handler: MsgPrivateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/msg/private/history",
				Handler: MsgPrivateHistoryHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/msg/recentcontact",
				Handler: MsgRecentcontactHandler(serverCtx),
			},
		},
	)
}

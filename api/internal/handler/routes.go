// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"shorturl-v2/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/shorten",
				Handler: GetShortURLHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/geturl",
				Handler: GetOriginalURLHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:alias",
				Handler: RedirectHandler(serverCtx),
			},
		},
	)
}

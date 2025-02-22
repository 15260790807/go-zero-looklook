// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	testgroup "looklook/app/test/cmd/api/internal/handler/testgroup"
	"looklook/app/test/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/testac/testac",
				Handler: testgroup.TestHandler(serverCtx),
			},
		},
		rest.WithPrefix("/test/v1"),
	)
}

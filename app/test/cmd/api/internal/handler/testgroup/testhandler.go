package testgroup

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/test/cmd/api/internal/logic/testgroup"
	"looklook/app/test/cmd/api/internal/svc"
	"looklook/app/test/cmd/api/internal/types"
)

func TestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := testgroup.NewTestLogic(r.Context(), svcCtx)
		resp, err := l.Test(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

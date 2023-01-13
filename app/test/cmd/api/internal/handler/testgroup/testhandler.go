package testgroup

import (
	"net/http"

	"looklook/app/test/cmd/api/internal/logic/testgroup"
	"looklook/app/test/cmd/api/internal/svc"
	"looklook/app/test/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func TestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TsReq
		if err := httpx.Parse(r, &req); err != nil {
			var resp types.TsResp
			resp.Msg = err.Error()
			httpx.OkJson(w, resp)
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

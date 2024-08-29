package demo

import (
	"net/http"

	"github.com/Sion-L/go-zero-demo/gateway/internal/logic/demo"
	"github.com/Sion-L/go-zero-demo/gateway/internal/svc"
	"github.com/Sion-L/go-zero-demo/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DemoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := demo.NewDemoLogic(r.Context(), svcCtx)
		resp, err := l.Demo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

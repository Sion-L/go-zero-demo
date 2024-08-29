package shorturl

import (
	"net/http"

	"github.com/Sion-L/go-zero-demo/gateway/internal/logic/shorturl"
	"github.com/Sion-L/go-zero-demo/gateway/internal/svc"
	"github.com/Sion-L/go-zero-demo/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ExpandHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExpandReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := shorturl.NewExpandLogic(r.Context(), svcCtx)
		resp, err := l.Expand(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

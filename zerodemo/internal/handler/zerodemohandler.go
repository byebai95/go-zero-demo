package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerodemo/internal/logic"
	"zerodemo/internal/svc"
	"zerodemo/internal/types"
)

func ZerodemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewZerodemoLogic(r.Context(), svcCtx)
		resp, err := l.Zerodemo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

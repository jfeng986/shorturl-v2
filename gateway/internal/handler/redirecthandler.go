package handler

import (
	"log"
	"net/http"

	"shorturl-v2/gateway/internal/logic"
	"shorturl-v2/gateway/internal/svc"
	"shorturl-v2/gateway/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RedirectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RedirectRequest
		alias := r.URL.Path[1:]
		req.ShortURL = "http://127.0.0.1:30000/" + alias
		log.Println("req.ShortURL:", req.ShortURL)

		l := logic.NewRedirectLogic(r.Context(), svcCtx)
		resp, err := l.Redirect(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

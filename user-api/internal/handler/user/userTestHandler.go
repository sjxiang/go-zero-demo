package user

import (
	"net/http"

	"github.com/sjxiang/go-zero-demo/user-api/internal/logic/user"
	"github.com/sjxiang/go-zero-demo/user-api/internal/svc"
	"github.com/sjxiang/go-zero-demo/user-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserTestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserTestReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserTestLogic(r.Context(), svcCtx)
		resp, err := l.UserTest(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

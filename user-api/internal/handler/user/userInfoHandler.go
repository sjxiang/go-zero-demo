package user

import (
	"net/http"

	"github.com/sjxiang/go-zero-demo/user-api/internal/logic/user"
	"github.com/sjxiang/go-zero-demo/user-api/internal/svc"
	"github.com/sjxiang/go-zero-demo/user-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 1. 参数解析
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// 2. 业务逻辑
		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

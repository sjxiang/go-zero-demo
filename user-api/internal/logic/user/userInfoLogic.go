package user

import (
	"context"

	"github.com/sjxiang/go-zero-demo/user-api/internal/svc"
	"github.com/sjxiang/go-zero-demo/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	
	// l.svcCtx.UserModel.FindOne()
	
	m := map[int64]string {
		1 : "Jisoo",
		2 : "Irene",
	}
	
	nickname := "unknown"
	if name, ok := m[req.UserId]; ok {
		nickname = name
	}

	return &types.UserInfoResp{
		UserId: req.UserId,
		NickName: nickname,
	}, nil
	
}

package user

import (
	"context"
	"errors"

	"github.com/sjxiang/go-zero-demo/user-api/internal/svc"
	"github.com/sjxiang/go-zero-demo/user-api/internal/types"
	"github.com/sjxiang/go-zero-demo/user-api/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateLogic) UserCreate(req *types.UserCreateReq) (resp *types.UserCreateResp, err error) {
	// todo: add your logic here and delete this line
	if err := l.svcCtx.UserModel.TranCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
	
		user := &model.User{}
		user.Mobile = req.Mobile
		user.Nickname = req.NickName

		// 添加 user
		dbResult, err := l.svcCtx.UserModel.TransInsert(ctx, session, user)  // 事务，坑有点大
		if err != nil {
			return err
		}
		userId, _ := dbResult.LastInsertId()

		// 添加 userData
		userData := &model.UserData{}
		userData.UserId = userId
		userData.Data = "xxxx"
		if _, err := l.svcCtx.UserDataModel.Insert(ctx, userData); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, errors.New("创建用户失败")
	}

	return &types.UserCreateResp{
		Flag: true,
	}, nil
}

package svc

import (
	"github.com/sjxiang/go-zero-demo/user-api/internal/config"
	"github.com/sjxiang/go-zero-demo/user-api/internal/middleware"
	"github.com/sjxiang/go-zero-demo/user-api/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      model.UserModel
	UserDataModel  model.UserDataModel
	TestMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserDataModel: model.NewUserDataModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		TestMiddleware: middleware.NewTestMiddleware().Handle,
	}
}

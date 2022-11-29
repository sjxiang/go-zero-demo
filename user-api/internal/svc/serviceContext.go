package svc

import (
	"github.com/sjxiang/go-zero-demo/user-api/internal/config"
	"github.com/sjxiang/go-zero-demo/user-api/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	UserModel model.UserModel
	// Kafka、Redis、UserModel 等
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
		// Todo ...
	}
}

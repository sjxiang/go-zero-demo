package svc

import (
	"github.com/sjxiang/go-zero-demo/user-api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	// Kafka、Redis、UserModel 等
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		// Todo ...
	}
}

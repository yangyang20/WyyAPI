package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"music/service/comment/api/internal/config"
	"music/service/comment/api/internal/middleware"
)

type ServiceContext struct {
	Config config.Config
	User   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   middleware.NewUserMiddleware().Handle,
	}
}

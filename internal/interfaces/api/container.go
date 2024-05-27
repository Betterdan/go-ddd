package api

import (
	"demo/internal/application/service"
	"demo/internal/infrastructure/config"
	"demo/internal/infrastructure/persistence"
    "demo/internal/domain/service" as domain_service

	"go.uber.org/dig"
)

func buildContainer(cfg *config.Config) *dig.Container {
	container := dig.New()

	// 注册配置
	container.Provide(func() *config.Config {
		return cfg
	})

	// 注册各模块依赖项
	persistence.Register(container)  //仓储持久层
	domain_service.Register(container)  //领域服务
	service.Register(container)  //应用服务
	Register(container)  
	return container
}

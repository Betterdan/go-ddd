package main

import (
	"demo/internal/infrastructure/cache"
	"demo/internal/infrastructure/config"
	"demo/internal/infrastructure/db"
	"demo/internal/infrastructure/logger"
	"demo/internal/interfaces/api"
	"demo/internal/interfaces/message"
	"go.uber.org/fx"
)

func main() {
	// 加载配置
	cfg, _ := config.LoadConfig()
	// beforeStartServer  load db mq redis log....
	//加载 log
	logger.InitLogger(cfg)
	//加载 redis pool
	cache.InitRedisCache(cfg)

	// 运行服务器
	fx.New(
		fx.Provide(func() *config.Config { return cfg }),
		fx.Provide(db.NewDB),
		api.Module,
		message.Module,
		fx.Invoke(
			api.StartServer,
			func(ms *message.MessageServer) {
				go ms.Start() // Start the message server
			}),
	).Run()
}

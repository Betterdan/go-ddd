package main

import (
	"demo/internal/infrastructure/cache"
	"demo/internal/infrastructure/config"
	"demo/internal/interfaces/api"
	"log"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	// beforeStartServer  load db mq redis log....

	//加载 redis pool
	cache.InitRedisCache(cfg)
	//加载 log

	//todo 加载 mq

	container := api.BuildContainer(cfg)

	// 运行服务器
	err = container.Invoke(func(server *api.Server) {
		server.Run()
	})
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}
}

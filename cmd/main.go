package main

import (
	"demo/internal/infrastructure/config"
	"demo/internal/interfaces/api"
	"log"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	// beforeStartServer  load db mq redis log....

	container := api.BuildContainer(cfg)

	// 运行服务器
	err = container.Invoke(func(server *api.Server) {
		server.Run()
	})
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}
}

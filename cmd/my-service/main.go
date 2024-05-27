package main

import (
	"demo/internal/infrastructure/config"
	"demo/internal/interfaces/api"
	"log"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// 创建并启动服务器
	err = api.StartServer(cfg)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

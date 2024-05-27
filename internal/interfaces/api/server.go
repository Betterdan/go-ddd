package api

import (
	"demo/internal/infrastructure/config"
	"log"
	"net/http"
)

type Server struct {
	cfg    *config.Config
	router *mux.Router
}

func NewServer(cfg *config.Config, router *mux.Router) *Server {
	return &Server{cfg: cfg, router: router}
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.cfg.ServerAddress, s.router)
}

func StartServer(cfg *config.Config) error {
	container := buildContainer(cfg)
	return container.Invoke(func(server *Server) {
		if err := server.Start(); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	})
}

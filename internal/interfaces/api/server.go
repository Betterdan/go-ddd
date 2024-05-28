package api

import (
	"context"
	"database/sql"
	"demo/internal/application/service"
	"demo/internal/infrastructure/config"
	"demo/internal/infrastructure/db"
	"demo/internal/infrastructure/persistence"
	"demo/internal/interfaces/api/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	Config     *config.Config
	Router     *gin.Engine
	httpServer *http.Server
	DB         *sql.DB
}

func NewServer(config *config.Config, router *gin.Engine, db *sql.DB) *Server {
	return &Server{
		Config: config,
		Router: router,
		DB:     db,
		httpServer: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port),
			Handler: router,
		},
	}
}

func (s *Server) Run() {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("could not listen on %s: %v\n", s.Config.Server.Host, err)
		}
	}()
	log.Printf("Server is running on %s\n", s.Config.Server.Host)

	s.gracefulShutdown()
}

func (s *Server) gracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	if err := s.DB.Close(); err != nil {
		log.Fatalf("Database forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}

func BuildContainer(allConfig *config.Config) *dig.Container {
	container := dig.New()

	container.Provide(func() *config.Config {
		return allConfig
	})
	container.Provide(db.NewDB)
	container.Provide(persistence.NewUserRepository)
	container.Provide(service.NewUserService)
	container.Provide(handler.NewUserHandler)

	container.Provide(NewHandlerList)
	container.Provide(NewRouterConfig)
	container.Provide(NewRouter)
	container.Provide(NewServer)

	return container
}

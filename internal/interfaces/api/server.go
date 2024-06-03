package api

import (
	"context"
	"database/sql"
	"demo/internal/application/service"
	service2 "demo/internal/domain/service"
	"demo/internal/infrastructure/config"
	"demo/internal/infrastructure/persistence"
	"demo/internal/interfaces/api/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Server represents the HTTP server
type Server struct {
	config     *config.Config
	router     *gin.Engine
	db         *sql.DB
	httpServer *http.Server
}

func NewServer(allConfig *config.Config, db *sql.DB, router *gin.Engine) *Server {
	return &Server{
		config: allConfig,
		db:     db,
		httpServer: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", allConfig.Server.Host, allConfig.Server.Port),
			Handler: router,
		},
	}
}

func StartServer(lc fx.Lifecycle, s *Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go s.Run()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			go s.GracefulShutdown()
			return nil
		},
	})
}

func (s *Server) Run() {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("could not listen on %s: %v\n", s.config.Server.Host, err)
		}
	}()
	log.Printf("Server is running on %s\n", s.config.Server.Host)
}

func (s *Server) GracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	if err := s.db.Close(); err != nil {
		log.Fatalf("Database forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}

var Module = fx.Options(
	fx.Provide(persistence.NewUserRepository),
	fx.Provide(service2.NewUserDomainService),
	fx.Provide(service.NewUserService),
	fx.Provide(handler.NewUserHandler),
	fx.Provide(NewHandlerList),
	fx.Provide(NewRouterConfig),
	fx.Provide(NewRouter),
	fx.Provide(NewServer),
)

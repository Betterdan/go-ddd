package api

import (
	"context"
	"demo/internal/application/service"
	service2 "demo/internal/domain/service"
	"demo/internal/infrastructure/config"
	"demo/internal/infrastructure/logger"
	"demo/internal/infrastructure/persistence"
	"demo/internal/interfaces/api/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Server represents the HTTP server
type Server struct {
	config     *config.Config
	db         *gorm.DB
	httpServer *http.Server
}

func NewServer(allConfig *config.Config, db *gorm.DB, router *gin.Engine) *Server {
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
			s.Run()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping server...")
			go s.GracefulShutdown()
			return nil
		},
	})
}

func (s *Server) Run() {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(fmt.Sprintf("could not listen on %s: %v\n", s.config.Server.Host, err))
		}
	}()
	logger.Info(fmt.Sprintf("Server is running on %s:%s\n", s.config.Server.Host, s.config.Server.Port))
}

func (s *Server) GracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	logger.Info(fmt.Sprintf("Shutting down server..."))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		logger.Fatal(fmt.Sprintf("Server forced to shutdown: %v", err))
	}

	sqlDb, err := s.db.DB()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to get underlying sql.DB from Gorm DB: %v", err))
	}
	if err := sqlDb.Close(); err != nil {
		logger.Fatal(fmt.Sprintf("Database forced to shutdown: %v", err))
	}

	logger.Info("Server exiting")
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

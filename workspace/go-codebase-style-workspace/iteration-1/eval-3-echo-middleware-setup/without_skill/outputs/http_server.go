package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"github.com/blkst8/example-service/internal/config"
	"github.com/blkst8/example-service/internal/http/handlers"
	"github.com/blkst8/example-service/internal/http/middlewares"
)

type Server struct {
	echo   *echo.Echo
	logger *zap.Logger
	config *config.Config
}

func NewServer(cfg *config.Config, logger *zap.Logger, h *handlers.Handler) *Server {
	e := echo.New()
	e.HideBanner = true

	// Register global middleware
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middlewares.ZapRequestLogger(logger, []string{"/healthz", "/metrics"}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: cfg.Server.AllowedOrigins,
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	// Public routes
	e.GET("/healthz", h.Healthz)
	e.GET("/metrics", h.Metrics)

	// Protected API routes
	api := e.Group("/api/v1")
	api.Use(middlewares.JWTMiddleware(cfg.Auth.JWTSecret))

	api.POST("/clients", h.CreateClient)
	api.GET("/clients/:id", h.GetClient)
	api.PUT("/clients/:id", h.UpdateClient)
	api.DELETE("/clients/:id", h.DeleteClient)

	return &Server{
		echo:   e,
		logger: logger,
		config: cfg,
	}
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.config.Server.Host, s.config.Server.Port)
	s.logger.Info("starting http server", zap.String("addr", addr))

	srv := &http.Server{
		Addr:         addr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := s.echo.StartServer(srv); err != nil && err != http.ErrServerClosed {
			s.logger.Error("server error", zap.Error(err))
		}
	}()

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("shutting down http server")
	return s.echo.Shutdown(ctx)
}

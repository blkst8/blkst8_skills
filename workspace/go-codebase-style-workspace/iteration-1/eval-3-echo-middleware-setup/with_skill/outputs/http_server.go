package httpserver

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"github.com/blkst8/example-service/internal/app"
	"github.com/blkst8/example-service/internal/config"
	"github.com/blkst8/example-service/internal/http/handlers"
	"github.com/blkst8/example-service/internal/http/middlewares"
	"github.com/blkst8/example-service/internal/log"
)

// Server wraps the Echo instance.
type Server struct {
	echo *echo.Echo
}

// NewServer constructs the Echo server, registers middleware, and wires all routes.
func NewServer(svc *app.Service) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// Global middleware — applied to every request.
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middlewares.ZapLogger(log.Logger, "/healthz", "/metrics"))
	e.Use(middleware.CORS())

	h := handlers.New(svc)

	// Unauthenticated endpoints.
	e.GET("/healthz", handlers.Healthz)
	e.GET("/metrics", handlers.Metrics)

	// Authenticated API routes — JWT middleware scoped to this group only.
	api := e.Group("/api/v1")
	api.Use(middlewares.JWTAuthentication())
	{
		api.POST("/clients", h.CreateClient)
		api.GET("/clients/:id", h.GetClient)
		api.PUT("/clients/:id", h.UpdateClient)
		api.DELETE("/clients/:id", h.DeleteClient)
	}

	return &Server{echo: e}
}

// Serve starts the HTTP server in a background goroutine.
func (s *Server) Serve() {
	cfg := config.C.HTTPServer

	srv := &http.Server{
		Addr:              cfg.Listen,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}

	go func() {
		if err := s.echo.StartServer(srv); err != nil && err != http.ErrServerClosed {
			log.Logger.Fatal("failed to start server", zap.Error(err))
		}
	}()
}

// Shutdown gracefully stops the HTTP server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

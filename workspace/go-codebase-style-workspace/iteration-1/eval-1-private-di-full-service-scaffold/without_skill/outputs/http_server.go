package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/example/invoice-service/internal/handler"
)

// Server wraps an Echo instance.
type Server struct {
	echo *echo.Echo
}

// New creates a new Server wired with the given handlers.
func New(invoiceHandler *handler.InvoiceHandler) *Server {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api/v1")
	invoiceHandler.Register(api)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	return &Server{echo: e}
}

// Start begins listening on the given address.
func (s *Server) Start(addr string) error {
	return s.echo.Start(addr)
}

// Shutdown gracefully stops the server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

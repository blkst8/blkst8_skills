package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// ZapRequestLogger returns an Echo middleware that logs incoming HTTP requests.
// The skipPaths parameter is a list of paths that should not be logged (e.g. /healthz, /metrics).
func ZapRequestLogger(logger *zap.Logger, skipPaths []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			path := req.URL.Path

			// Check if this path should be skipped
			for _, skipPath := range skipPaths {
				if path == skipPath {
					return next(c)
				}
			}

			start := time.Now()
			err := next(c)
			latency := time.Since(start)

			res := c.Response()

			logger.Info("http request",
				zap.String("method", req.Method),
				zap.String("path", path),
				zap.String("query", req.URL.RawQuery),
				zap.Int("status", res.Status),
				zap.Duration("latency", latency),
				zap.String("remote_ip", c.RealIP()),
				zap.String("request_id", res.Header().Get(echo.HeaderXRequestID)),
			)

			return err
		}
	}
}

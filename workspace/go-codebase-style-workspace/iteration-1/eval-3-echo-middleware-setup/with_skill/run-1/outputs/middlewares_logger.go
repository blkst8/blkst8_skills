package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// ZapLogger returns an Echo middleware that logs each HTTP request using the
// provided *zap.Logger. Paths in skipURLs are skipped entirely — useful for
// /healthz and /metrics endpoints that would otherwise flood production logs.
func ZapLogger(logger *zap.Logger, skipURLs ...string) echo.MiddlewareFunc {
	skipSet := make(map[string]struct{}, len(skipURLs))
	for _, u := range skipURLs {
		skipSet[u] = struct{}{}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)

			if _, skip := skipSet[c.Request().URL.Path]; skip {
				return err
			}

			logger.Info("request",
				zap.String("method", c.Request().Method),
				zap.String("uri", c.Request().RequestURI),
				zap.Int("status", c.Response().Status),
				zap.Duration("latency", time.Since(start)),
			)

			return err
		}
	}
}

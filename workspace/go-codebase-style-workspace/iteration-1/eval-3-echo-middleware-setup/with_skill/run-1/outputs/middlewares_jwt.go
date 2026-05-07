package middlewares

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/blkst8/example-service/internal/config"
	"github.com/blkst8/example-service/internal/log"
)

// Claims holds the JWT registered claims and application-specific fields.
type Claims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

// JWTAuthentication returns an Echo middleware that validates Bearer tokens
// and stores the parsed *Claims in the request context under the key "user".
func JWTAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.ErrUnauthorized
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
				return echo.ErrUnauthorized
			}

			claims, err := validateToken(parts[1], []byte(config.C.Auth.JWTSecret))
			if err != nil {
				log.Logger.Warn("invalid jwt token", zap.Error(err))
				return echo.ErrUnauthorized
			}

			c.Set("user", claims)

			return next(c)
		}
	}
}

func validateToken(raw string, secret []byte) (*Claims, error) {
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(raw, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse jwt: %w", err)
	}

	return claims, nil
}

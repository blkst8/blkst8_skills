package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/blkst8/client-service/internal/app"
	"github.com/blkst8/client-service/internal/log"
	"github.com/blkst8/client-service/internal/repository"
)

// GetClient handles GET /api/v1/clients/:id.
func GetClient(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid client id")
	}

	client, err := app.A.Service.Client.Get(ctx.Request().Context(), uint32(id))
	if err != nil {
		if errors.Is(err, repository.ErrClientNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "client not found")
		}
		log.Logger.Error("failed to get client", zap.Uint64("id", id), zap.Error(err))
		return err
	}

	return ctx.JSON(http.StatusOK, client)
}
